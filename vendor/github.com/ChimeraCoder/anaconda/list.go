package anaconda

type ListResponse struct {
	PreviousCursor int    `json:"previous_cursor"`
	NextCursor     int    `json:"next_cursor"`
	Lists          []List `json:"lists"`
}

type AddUserToListResponse struct {
	Users []User `json:"users"`
}

type List struct {
	Slug            string `json:"slug"`
	Name            string `json:"name"`
	URL             string `json:"uri"`
	CreatedAt       string `json:"created_at"`
	Id              int64  `json:"id"`
	SubscriberCount int64  `json:"subscriber_count"`
	MemberCount     int64  `json:"member_count"`
	Mode            string `json:"mode"`
	FullName        string `json:"full_name"`
	Description     string `json:"description"`
	User            User   `json:"user"`
	Following       bool   `json:"following"`
}
