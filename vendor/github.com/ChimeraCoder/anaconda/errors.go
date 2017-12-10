package anaconda

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	//Error code defintions match the Twitter documentation
	//https://dev.twitter.com/docs/error-codes-responses
	TwitterErrorCouldNotAuthenticate    = 32
	TwitterErrorDoesNotExist            = 34
	TwitterErrorAccountSuspended        = 64
	TwitterErrorApi1Deprecation         = 68 //This should never be needed
	TwitterErrorRateLimitExceeded       = 88
	TwitterErrorInvalidToken            = 89
	TwitterErrorOverCapacity            = 130
	TwitterErrorInternalError           = 131
	TwitterErrorCouldNotAuthenticateYou = 135
	TwitterErrorStatusIsADuplicate      = 187
	TwitterErrorBadAuthenticationData   = 215
	TwitterErrorUserMustVerifyLogin     = 231

	// Undocumented by Twitter, but may be returned instead of 34
	TwitterErrorDoesNotExist2 = 144
)

type ApiError struct {
	StatusCode int
	Header     http.Header
	Body       string
	Decoded    TwitterErrorResponse
	URL        *url.URL
}

func newApiError(resp *http.Response) *ApiError {
	// TODO don't ignore this error
	// TODO don't use ReadAll
	p, _ := ioutil.ReadAll(resp.Body)

	var twitterErrorResp TwitterErrorResponse
	_ = json.Unmarshal(p, &twitterErrorResp)
	return &ApiError{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(p),
		Decoded:    twitterErrorResp,
		URL:        resp.Request.URL,
	}
}

// ApiError supports the error interface
func (aerr ApiError) Error() string {
	return fmt.Sprintf("Get %s returned status %d, %s", aerr.URL, aerr.StatusCode, aerr.Body)
}

// Check to see if an error is a Rate Limiting error. If so, find the next available window in the header.
// Use like so:
//
//    if aerr, ok := err.(*ApiError); ok {
//  	  if isRateLimitError, nextWindow := aerr.RateLimitCheck(); isRateLimitError {
//       	<-time.After(nextWindow.Sub(time.Now()))
//  	  }
//    }
//
func (aerr *ApiError) RateLimitCheck() (isRateLimitError bool, nextWindow time.Time) {
	// TODO  check for error code 130, which also signifies a rate limit
	if aerr.StatusCode == 429 {
		if reset := aerr.Header.Get("X-Rate-Limit-Reset"); reset != "" {
			if resetUnix, err := strconv.ParseInt(reset, 10, 64); err == nil {
				resetTime := time.Unix(resetUnix, 0)
				// Reject any time greater than an hour away
				if resetTime.Sub(time.Now()) > time.Hour {
					return true, time.Now().Add(15 * time.Minute)
				}

				return true, resetTime
			}
		}
	}

	return false, time.Time{}
}

//TwitterErrorResponse has an array of Twitter error messages
//It satisfies the "error" interface
//For the most part, Twitter seems to return only a single error message
//Currently, we assume that this always contains exactly one error message
type TwitterErrorResponse struct {
	Errors []TwitterError `json:"errors"`
}

func (tr TwitterErrorResponse) First() error {
	return tr.Errors[0]
}

func (tr TwitterErrorResponse) Error() string {
	return tr.Errors[0].Message
}

//TwitterError represents a single Twitter error messages/code pair
type TwitterError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (te TwitterError) Error() string {
	return te.Message
}
