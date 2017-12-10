package anaconda

import (
	"net/url"
	"strconv"
)

func (a TwitterApi) GetBlocksList(v url.Values) (c UserCursor, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/blocks/list.json", v, &c, _GET, response_ch}
	return c, (<-response_ch).err
}

func (a TwitterApi) GetBlocksIds(v url.Values) (c Cursor, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/blocks/ids.json", v, &c, _GET, response_ch}
	return c, (<-response_ch).err
}

func (a TwitterApi) BlockUser(screenName string, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("screen_name", screenName)
	return a.Block(v)
}

func (a TwitterApi) BlockUserId(id int64, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(id, 10))
	return a.Block(v)
}

func (a TwitterApi) Block(v url.Values) (user User, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/blocks/create.json", v, &user, _POST, response_ch}
	return user, (<-response_ch).err
}

func (a TwitterApi) UnblockUser(screenName string, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("screen_name", screenName)
	return a.Unblock(v)
}

func (a TwitterApi) UnblockUserId(id int64, v url.Values) (user User, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(id, 10))
	return a.Unblock(v)
}

func (a TwitterApi) Unblock(v url.Values) (user User, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/blocks/destroy.json", v, &user, _POST, response_ch}
	return user, (<-response_ch).err
}
