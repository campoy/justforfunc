package anaconda

import (
	"net/url"
	"strconv"
)

func (a TwitterApi) GetDirectMessages(v url.Values) (messages []DirectMessage, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/direct_messages.json", v, &messages, _GET, response_ch}
	return messages, (<-response_ch).err
}

func (a TwitterApi) GetDirectMessagesSent(v url.Values) (messages []DirectMessage, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/direct_messages/sent.json", v, &messages, _GET, response_ch}
	return messages, (<-response_ch).err
}

func (a TwitterApi) GetDirectMessagesShow(v url.Values) (messages []DirectMessage, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/direct_messages/show.json", v, &messages, _GET, response_ch}
	return messages, (<-response_ch).err
}

// https://dev.twitter.com/docs/api/1.1/post/direct_messages/new
func (a TwitterApi) PostDMToScreenName(text, screenName string) (message DirectMessage, err error) {
	v := url.Values{}
	v.Set("screen_name", screenName)
	v.Set("text", text)
	return a.postDirectMessagesImpl(v)
}

// https://dev.twitter.com/docs/api/1.1/post/direct_messages/new
func (a TwitterApi) PostDMToUserId(text string, userId int64) (message DirectMessage, err error) {
	v := url.Values{}
	v.Set("user_id", strconv.FormatInt(userId, 10))
	v.Set("text", text)
	return a.postDirectMessagesImpl(v)
}

// DeleteDirectMessage will destroy (delete) the direct message with the specified ID.
// https://dev.twitter.com/rest/reference/post/direct_messages/destroy
func (a TwitterApi) DeleteDirectMessage(id int64, includeEntities bool) (message DirectMessage, err error) {
	v := url.Values{}
	v.Set("id", strconv.FormatInt(id, 10))
	v.Set("include_entities", strconv.FormatBool(includeEntities))
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/direct_messages/destroy.json", v, &message, _POST, response_ch}
	return message, (<-response_ch).err
}

func (a TwitterApi) postDirectMessagesImpl(v url.Values) (message DirectMessage, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/direct_messages/new.json", v, &message, _POST, response_ch}
	return message, (<-response_ch).err
}
