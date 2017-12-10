package anaconda

import (
	"net/url"
	"strconv"
)

type Media struct {
	MediaID       int64  `json:"media_id"`
	MediaIDString string `json:"media_id_string"`
	Size          int    `json:"size"`
	Image         Image  `json:"image"`
}

type Image struct {
	W         int    `json:"w"`
	H         int    `json:"h"`
	ImageType string `json:"image_type"`
}

type ChunkedMedia struct {
	MediaID          int64  `json:"media_id"`
	MediaIDString    string `json:"media_id_string"`
	ExpiresAfterSecs int    `json:"expires_after_secs"`
}

type Video struct {
	VideoType string `json:"video_type"`
}

type VideoMedia struct {
	MediaID          int64  `json:"media_id"`
	MediaIDString    string `json:"media_id_string"`
	Size             int    `json:"size"`
	ExpiresAfterSecs int    `json:"expires_after_secs"`
	Video            Video  `json:"video"`
}

func (a TwitterApi) UploadMedia(base64String string) (media Media, err error) {
	v := url.Values{}
	v.Set("media_data", base64String)

	var mediaResponse Media

	response_ch := make(chan response)
	a.queryQueue <- query{UploadBaseUrl + "/media/upload.json", v, &mediaResponse, _POST, response_ch}
	return mediaResponse, (<-response_ch).err
}

func (a TwitterApi) UploadVideoInit(totalBytes int, mimeType string) (chunkedMedia ChunkedMedia, err error) {
	v := url.Values{}
	v.Set("command", "INIT")
	v.Set("media_type", mimeType)
	v.Set("total_bytes", strconv.FormatInt(int64(totalBytes), 10))

	var mediaResponse ChunkedMedia

	response_ch := make(chan response)
	a.queryQueue <- query{UploadBaseUrl + "/media/upload.json", v, &mediaResponse, _POST, response_ch}
	return mediaResponse, (<-response_ch).err
}

func (a TwitterApi) UploadVideoAppend(mediaIdString string,
	segmentIndex int, base64String string) error {

	v := url.Values{}
	v.Set("command", "APPEND")
	v.Set("media_id", mediaIdString)
	v.Set("media_data", base64String)
	v.Set("segment_index", strconv.FormatInt(int64(segmentIndex), 10))

	var emptyResponse interface{}

	response_ch := make(chan response)
	a.queryQueue <- query{UploadBaseUrl + "/media/upload.json", v, &emptyResponse, _POST, response_ch}
	return (<-response_ch).err
}

func (a TwitterApi) UploadVideoFinalize(mediaIdString string) (videoMedia VideoMedia, err error) {
	v := url.Values{}
	v.Set("command", "FINALIZE")
	v.Set("media_id", mediaIdString)

	var mediaResponse VideoMedia

	response_ch := make(chan response)
	a.queryQueue <- query{UploadBaseUrl + "/media/upload.json", v, &mediaResponse, _POST, response_ch}
	return mediaResponse, (<-response_ch).err
}
