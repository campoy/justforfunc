package anaconda

import (
	"net/url"
)

type Configuration struct {
	CharactersReservedPerMedia int      `json:"characters_reserved_per_media"`
	MaxMediaPerUpload          int      `json:"max_media_per_upload"`
	NonUsernamePaths           []string `json:"non_username_paths"`
	PhotoSizeLimit             int      `json:"photo_size_limit"`
	PhotoSizes                 struct {
		Thumb  photoSize `json:"thumb"`
		Small  photoSize `json:"small"`
		Medium photoSize `json:"medium"`
		Large  photoSize `json:"large"`
	} `json:"photo_sizes"`
	ShortUrlLength      int `json:"short_url_length"`
	ShortUrlLengthHttps int `json:"short_url_length_https"`
}

type photoSize struct {
	H      int    `json:"h"`
	W      int    `json:"w"`
	Resize string `json:"resize"`
}

func (a TwitterApi) GetConfiguration(v url.Values) (conf Configuration, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/help/configuration.json", v, &conf, _GET, response_ch}
	return conf, (<-response_ch).err
}
