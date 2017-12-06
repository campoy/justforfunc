package anaconda

import (
	"net/url"
	"strconv"
)

func (a TwitterApi) GetMutedUsersList(v url.Values) (c UserCursor, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/mutes/users/list.json", v, &c, _GET, response_ch}
	return c, (<-response_ch).err
}

func (a TwitterApi) GetMutedUsersIds(v url.Values) (c Cursor, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/mutes/users/ids.json", v, &c, _GET, response_ch}
	return c, (<-response_ch).err
}

func (a TwitterApi) MuteUser(screenName string, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("screen_name", screenName)
	return a.Mute(v)
}

func (a TwitterApi) MuteUserId(id int64, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(id, 10))
	return a.Mute(v)
}

func (a TwitterApi) Mute(v url.Values) (user User, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/mutes/users/create.json", v, &user, _POST, response_ch}
	return user, (<-response_ch).err
}

func (a TwitterApi) UnmuteUser(screenName string, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("screen_name", screenName)
	return a.Unmute(v)
}

func (a TwitterApi) UnmuteUserId(id int64, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(id, 10))
	return a.Unmute(v)
}

func (a TwitterApi) Unmute(v url.Values) (user User, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/mutes/users/destroy.json", v, &user, _POST, response_ch}
	return user, (<-response_ch).err
}
