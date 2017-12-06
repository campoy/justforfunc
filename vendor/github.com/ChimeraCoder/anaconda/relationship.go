package anaconda

import (
	"net/url"
)

type RelationshipResponse struct {
	Relationship Relationship `json:"relationship"`
}
type Relationship struct {
	Target Target `json:"target"`
	Source Source `json:"source"`
}
type Target struct {
	Id          int64  `json:"id"`
	Id_str      string `json:"id_str"`
	Screen_name string `json:"screen_name"`
	Following   bool   `json:"following"`
	Followed_by bool   `json:"followed_by"`
}
type Source struct {
	Id                    int64
	Id_str                string
	Screen_name           string
	Following             bool
	Followed_by           bool
	Can_dm                bool
	Blocking              bool
	Muting                bool
	Marked_spam           bool
	All_replies           bool
	Want_retweets         bool
	Notifications_enabled bool
}

func (a TwitterApi) GetFriendshipsShow(v url.Values) (relationshipResponse RelationshipResponse, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/friendships/show.json", v, &relationshipResponse, _GET, response_ch}
	return relationshipResponse, (<-response_ch).err
}
