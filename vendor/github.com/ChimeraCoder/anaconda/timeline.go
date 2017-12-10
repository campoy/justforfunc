package anaconda

import (
	"net/url"
)

// GetHomeTimeline returns the most recent tweets and retweets posted by the user
// and the users that they follow.
// https://dev.twitter.com/docs/api/1.1/get/statuses/home_timeline
// By default, include_entities is set to "true"
func (a TwitterApi) GetHomeTimeline(v url.Values) (timeline []Tweet, err error) {
	if v == nil {
		v = url.Values{}
	}

	if val := v.Get("include_entities"); val == "" {
		v.Set("include_entities", "true")
	}

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/statuses/home_timeline.json", v, &timeline, _GET, response_ch}
	return timeline, (<-response_ch).err
}

func (a TwitterApi) GetUserTimeline(v url.Values) (timeline []Tweet, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/statuses/user_timeline.json", v, &timeline, _GET, response_ch}
	return timeline, (<-response_ch).err
}

func (a TwitterApi) GetMentionsTimeline(v url.Values) (timeline []Tweet, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/statuses/mentions_timeline.json", v, &timeline, _GET, response_ch}
	return timeline, (<-response_ch).err
}

func (a TwitterApi) GetRetweetsOfMe(v url.Values) (tweets []Tweet, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/statuses/retweets_of_me.json", v, &tweets, _GET, response_ch}
	return tweets, (<-response_ch).err
}
