package anaconda

type User struct {
	ContributorsEnabled            bool     `json:"contributors_enabled"`
	CreatedAt                      string   `json:"created_at"`
	DefaultProfile                 bool     `json:"default_profile"`
	DefaultProfileImage            bool     `json:"default_profile_image"`
	Description                    string   `json:"description"`
	Entities                       Entities `json:"entities"`
	FavouritesCount                int      `json:"favourites_count"`
	FollowRequestSent              bool     `json:"follow_request_sent"`
	FollowersCount                 int      `json:"followers_count"`
	Following                      bool     `json:"following"`
	FriendsCount                   int      `json:"friends_count"`
	GeoEnabled                     bool     `json:"geo_enabled"`
	HasExtendedProfile             bool     `json:"has_extended_profile"`
	Id                             int64    `json:"id"`
	IdStr                          string   `json:"id_str"`
	IsTranslator                   bool     `json:"is_translator"`
	IsTranslationEnabled           bool     `json:"is_translation_enabled"`
	Lang                           string   `json:"lang"` // BCP-47 code of user defined language
	ListedCount                    int64    `json:"listed_count"`
	Location                       string   `json:"location"` // User defined location
	Name                           string   `json:"name"`
	Notifications                  bool     `json:"notifications"`
	ProfileBackgroundColor         string   `json:"profile_background_color"`
	ProfileBackgroundImageURL      string   `json:"profile_background_image_url"`
	ProfileBackgroundImageUrlHttps string   `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool     `json:"profile_background_tile"`
	ProfileBannerURL               string   `json:"profile_banner_url"`
	ProfileImageURL                string   `json:"profile_image_url"`
	ProfileImageUrlHttps           string   `json:"profile_image_url_https"`
	ProfileLinkColor               string   `json:"profile_link_color"`
	ProfileSidebarBorderColor      string   `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string   `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string   `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool     `json:"profile_use_background_image"`
	Protected                      bool     `json:"protected"`
	ScreenName                     string   `json:"screen_name"`
	ShowAllInlineMedia             bool     `json:"show_all_inline_media"`
	Status                         *Tweet   `json:"status"` // Only included if the user is a friend
	StatusesCount                  int64    `json:"statuses_count"`
	TimeZone                       string   `json:"time_zone"`
	URL                            string   `json:"url"`
	UtcOffset                      int      `json:"utc_offset"`
	Verified                       bool     `json:"verified"`
	WithheldInCountries            []string `json:"withheld_in_countries"`
	WithheldScope                  string   `json:"withheld_scope"`
}

// Provide language translator from BCP-47 to human readable format for Lang field?
// Available through golang.org/x/text/language, deserves further investigation
