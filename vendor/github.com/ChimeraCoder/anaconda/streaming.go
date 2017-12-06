package anaconda

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/dustin/go-jsonpointer"
)

const (
	BaseUrlUserStream = "https://userstream.twitter.com/1.1"
	BaseUrlSiteStream = "https://sitestream.twitter.com/1.1"
	BaseUrlStream     = "https://stream.twitter.com/1.1"
)

// messages

type StatusDeletionNotice struct {
	Id        int64  `json:"id"`
	IdStr     string `json:"id_str"`
	UserId    int64  `json:"user_id"`
	UserIdStr string `json:"user_id_str"`
}
type statusDeletionNotice struct {
	Delete *struct {
		Status *StatusDeletionNotice `json:"status"`
	} `json:"delete"`
}

type LocationDeletionNotice struct {
	UserId          int64  `json:"user_id"`
	UserIdStr       string `json:"user_id_str"`
	UpToStatusId    int64  `json:"up_to_status_id"`
	UpToStatusIdStr string `json:"up_to_status_id_str"`
}
type locationDeletionNotice struct {
	ScrubGeo *LocationDeletionNotice `json:"scrub_geo"`
}

type LimitNotice struct {
	Track int64 `json:"track"`
}
type limitNotice struct {
	Limit *LimitNotice `json:"limit"`
}

type StatusWithheldNotice struct {
	Id                  int64    `json:"id"`
	UserId              int64    `json:"user_id"`
	WithheldInCountries []string `json:"withheld_in_countries"`
}
type statusWithheldNotice struct {
	StatusWithheld *StatusWithheldNotice `json:"status_withheld"`
}

type UserWithheldNotice struct {
	Id                  int64    `json:"id"`
	WithheldInCountries []string `json:"withheld_in_countries"`
}
type userWithheldNotice struct {
	UserWithheld *UserWithheldNotice `json:"user_withheld"`
}

type DisconnectMessage struct {
	Code       int64  `json:"code"`
	StreamName string `json:"stream_name"`
	Reason     string `json:"reason"`
}
type disconnectMessage struct {
	Disconnect *DisconnectMessage `json:"disconnect"`
}

type StallWarning struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	PercentFull int64  `json:"percent_full"`
}
type stallWarning struct {
	Warning *StallWarning `json:"warning"`
}

type FriendsList []int64
type friendsList struct {
	Friends *FriendsList `json:"friends"`
}

type streamDirectMessage struct {
	DirectMessage *DirectMessage `json:"direct_message"`
}

type Event struct {
	Target    *User  `json:"target"`
	Source    *User  `json:"source"`
	Event     string `json:"event"`
	CreatedAt string `json:"created_at"`
}

type EventList struct {
	Event
	TargetObject *List `json:"target_object"`
}

type EventTweet struct {
	Event
	TargetObject *Tweet `json:"target_object"`
}

type TooManyFollow struct {
	Warning *struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		UserId  int64  `json:"user_id"`
	} `json:"warning"`
}

// TODO: Site Stream messages. I cant test.

// Stream allows you to stream using one of the
// PublicStream* or UserStream api methods
//
// A go loop is started an gives you an stream that sends interface{}
// objects through it's chan C
// Objects which you can cast into a tweet like this :
//    t, ok := o.(twitter.Tweet) // try casting into a tweet
//    if !ok {
//      log.Debug("Recieved non tweet message")
//    }
//
// If we can't stream the chan will be closed.
// Otherwise the loop will connect and send streams in the chan.
// It will also try to reconnect itself after an exponential backoff time
// if the connection is lost
// If twitter response is one of 420, 429 or 503 (meaning "wait a sec")
// the loop retries to open the socket with a simple autogrowing backoff.
//
// When finished streaming call stream.Stop() to initiate termination process.
//

type Stream struct {
	api TwitterApi
	C   chan interface{}
	run bool
}

func (s *Stream) listen(response http.Response) {
	if response.Body != nil {
		defer response.Body.Close()
	}

	s.api.Log.Notice("Listenning to twitter socket")
	defer s.api.Log.Notice("twitter socket closed, leaving loop")

	scanner := bufio.NewScanner(response.Body)

	for scanner.Scan() && s.run {
		j := scanner.Bytes()
		if len(j) == 0 {
			s.api.Log.Debug("Empty bytes... Moving along")
		} else {
			s.C <- jsonToKnownType(j)
		}
	}
}

func jsonToKnownType(j []byte) interface{} {
	// TODO: DRY
	if o := new(Tweet); jsonAsStruct(j, "/source", &o) {
		return *o
	} else if o := new(statusDeletionNotice); jsonAsStruct(j, "/delete", &o) {
		return *o.Delete.Status
	} else if o := new(locationDeletionNotice); jsonAsStruct(j, "/scrub_geo", &o) {
		return *o.ScrubGeo
	} else if o := new(limitNotice); jsonAsStruct(j, "/limit", &o) {
		return *o.Limit
	} else if o := new(statusWithheldNotice); jsonAsStruct(j, "/status_withheld", &o) {
		return *o.StatusWithheld
	} else if o := new(userWithheldNotice); jsonAsStruct(j, "/user_withheld", &o) {
		return *o.UserWithheld
	} else if o := new(disconnectMessage); jsonAsStruct(j, "/disconnect", &o) {
		return *o.Disconnect
	} else if o := new(stallWarning); jsonAsStruct(j, "/warning", &o) {
		return *o.Warning
	} else if o := new(friendsList); jsonAsStruct(j, "/friends", &o) {
		return *o.Friends
	} else if o := new(streamDirectMessage); jsonAsStruct(j, "/direct_message", &o) {
		return *o.DirectMessage
	} else if o := new(EventTweet); jsonAsStruct(j, "/target_object/source", &o) {
		return *o
	} else if o := new(EventList); jsonAsStruct(j, "/target_object/slug", &o) {
		return *o
	} else if o := new(Event); jsonAsStruct(j, "/target_object", &o) {
		return *o
	} else {
		return nil
	}
}

func (s *Stream) requestStream(urlStr string, v url.Values, method int) (resp *http.Response, err error) {
	switch method {
	case _GET:
		return oauthClient.Get(s.api.HttpClient, s.api.Credentials, urlStr, v)
	case _POST:
		return oauthClient.Post(s.api.HttpClient, s.api.Credentials, urlStr, v)
	default:
	}
	return nil, fmt.Errorf("HTTP method not yet supported")
}

func (s *Stream) loop(urlStr string, v url.Values, method int) {
	defer s.api.Log.Debug("Leaving request stream loop")
	defer close(s.C)

	rlb := NewHTTP420ErrBackoff()
	for s.run {
		resp, err := s.requestStream(urlStr, v, method)
		if err != nil {
			if err == io.EOF {
				// Sometimes twitter closes the stream
				// right away with EOF as of a rate limit
				resp.StatusCode = 420
			} else {
				s.api.Log.Criticalf("Cannot request stream : %s", err)
				return
			}
		}
		s.api.Log.Debugf("Response status=%s code=%d", resp.Status, resp.StatusCode)

		switch resp.StatusCode {
		case 200, 304:
			s.listen(*resp)
			rlb.Reset()
		case 420, 429, 503:
			s.api.Log.Noticef("Twitter streaming: backing off as got : %+s", resp.Status)
			rlb.BackOff()
		case 400, 401, 403, 404, 406, 410, 422, 500, 502, 504:
			s.api.Log.Criticalf("Twitter streaming: leaving after an irremediable error: %+s", resp.Status)
			return
		default:
			s.api.Log.Notice("Received unknown status: %+s", resp.StatusCode)
		}

	}
}

func (s *Stream) Stop() {
	s.run = false
}

func (s *Stream) start(urlStr string, v url.Values, method int) {
	s.run = true
	go s.loop(urlStr, v, method)
}

func (a TwitterApi) newStream(urlStr string, v url.Values, method int) *Stream {
	stream := Stream{
		api: a,
		C:   make(chan interface{}),
	}

	stream.start(urlStr, v, method)
	return &stream
}

func (a TwitterApi) UserStream(v url.Values) (stream *Stream) {
	return a.newStream(BaseUrlUserStream+"/user.json", v, _GET)
}

func (a TwitterApi) PublicStreamSample(v url.Values) (stream *Stream) {
	return a.newStream(BaseUrlStream+"/statuses/sample.json", v, _GET)
}

// XXX: To use this API authority is requied. but I dont have this. I cant test.
func (a TwitterApi) PublicStreamFirehose(v url.Values) (stream *Stream) {
	return a.newStream(BaseUrlStream+"/statuses/firehose.json", v, _GET)
}

// XXX: PublicStream(Track|Follow|Locations) func is needed?
func (a TwitterApi) PublicStreamFilter(v url.Values) (stream *Stream) {
	return a.newStream(BaseUrlStream+"/statuses/filter.json", v, _POST)
}

// XXX: To use this API authority is requied. but I dont have this. I cant test.
func (a TwitterApi) SiteStream(v url.Values) (stream *Stream) {
	return a.newStream(BaseUrlSiteStream+"/site.json", v, _GET)
}

func jsonAsStruct(j []byte, path string, obj interface{}) (res bool) {
	if v, _ := jsonpointer.Find(j, path); v == nil {
		return false
	}
	err := json.Unmarshal(j, obj)
	return err == nil
}
