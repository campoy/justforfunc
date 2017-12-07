package anaconda

import (
	"fmt"
	"net/url"
	"strconv"
)

func (a TwitterApi) GetTweet(id int64, v url.Values) (tweet Tweet, err error) {
	v = cleanValues(v)
	v.Set("id", strconv.FormatInt(id, 10))

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/statuses/show.json", v, &tweet, _GET, response_ch}
	return tweet, (<-response_ch).err
}

func (a TwitterApi) GetTweetsLookupByIds(ids []int64, v url.Values) (tweet []Tweet, err error) {
	var pids string
	for w, i := range ids {
		pids += strconv.FormatInt(i, 10)
		if w != len(ids)-1 {
			pids += ","
		}
	}
	v = cleanValues(v)
	v.Set("id", pids)
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/statuses/lookup.json", v, &tweet, _GET, response_ch}
	return tweet, (<-response_ch).err
}

func (a TwitterApi) GetRetweets(id int64, v url.Values) (tweets []Tweet, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + fmt.Sprintf("/statuses/retweets/%d.json", id), v, &tweets, _GET, response_ch}
	return tweets, (<-response_ch).err
}

//PostTweet will create a tweet with the specified status message
func (a TwitterApi) PostTweet(status string, v url.Values) (tweet Tweet, err error) {
	v = cleanValues(v)
	v.Set("status", status)
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/statuses/update.json", v, &tweet, _POST, response_ch}
	return tweet, (<-response_ch).err
}

//DeleteTweet will destroy (delete) the status (tweet) with the specified ID, assuming that the authenticated user is the author of the status (tweet).
//If trimUser is set to true, only the user's Id will be provided in the user object returned.
func (a TwitterApi) DeleteTweet(id int64, trimUser bool) (tweet Tweet, err error) {
	v := url.Values{}
	if trimUser {
		v.Set("trim_user", "t")
	}
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + fmt.Sprintf("/statuses/destroy/%d.json", id), v, &tweet, _POST, response_ch}
	return tweet, (<-response_ch).err
}

//Retweet will retweet the status (tweet) with the specified ID.
//trimUser functions as in DeleteTweet
func (a TwitterApi) Retweet(id int64, trimUser bool) (rt Tweet, err error) {
	v := url.Values{}
	if trimUser {
		v.Set("trim_user", "t")
	}
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + fmt.Sprintf("/statuses/retweet/%d.json", id), v, &rt, _POST, response_ch}
	return rt, (<-response_ch).err
}

//UnRetweet will renove retweet Untweets a retweeted status.
//Returns the original Tweet with retweet details embedded.
//
//https://dev.twitter.com/rest/reference/post/statuses/unretweet/id
//trim_user: tweet returned in a timeline will include a user object
//including only the status authors numerical ID.
func (a TwitterApi) UnRetweet(id int64, trimUser bool) (rt Tweet, err error) {
	v := url.Values{}
	if trimUser {
		v.Set("trim_user", "t")
	}
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + fmt.Sprintf("/statuses/unretweet/%d.json", id), v, &rt, _POST, response_ch}
	return rt, (<-response_ch).err
}

// Favorite will favorite the status (tweet) with the specified ID.
// https://dev.twitter.com/docs/api/1.1/post/favorites/create
func (a TwitterApi) Favorite(id int64) (rt Tweet, err error) {
	v := url.Values{}
	v.Set("id", fmt.Sprint(id))
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + fmt.Sprintf("/favorites/create.json"), v, &rt, _POST, response_ch}
	return rt, (<-response_ch).err
}

// Un-favorites the status specified in the ID parameter as the authenticating user.
// Returns the un-favorited status in the requested format when successful.
// https://dev.twitter.com/docs/api/1.1/post/favorites/destroy
func (a TwitterApi) Unfavorite(id int64) (rt Tweet, err error) {
	v := url.Values{}
	v.Set("id", fmt.Sprint(id))
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + fmt.Sprintf("/favorites/destroy.json"), v, &rt, _POST, response_ch}
	return rt, (<-response_ch).err
}
