package anaconda

import (
	"encoding/json"
	"fmt"
	"time"
)

type Tweet struct {
	Contributors                []int64                `json:"contributors"`
	Coordinates                 *Coordinates           `json:"coordinates"`
	CreatedAt                   string                 `json:"created_at"`
	DisplayTextRange            []int                  `json:"display_text_range"`
	Entities                    Entities               `json:"entities"`
	ExtendedEntities            Entities               `json:"extended_entities"`
	ExtendedTweet               ExtendedTweet          `json:"extended_tweet"`
	FavoriteCount               int                    `json:"favorite_count"`
	Favorited                   bool                   `json:"favorited"`
	FilterLevel                 string                 `json:"filter_level"`
	FullText                    string                 `json:"full_text"`
	HasExtendedProfile          bool                   `json:"has_extended_profile"`
	Id                          int64                  `json:"id"`
	IdStr                       string                 `json:"id_str"`
	InReplyToScreenName         string                 `json:"in_reply_to_screen_name"`
	InReplyToStatusID           int64                  `json:"in_reply_to_status_id"`
	InReplyToStatusIdStr        string                 `json:"in_reply_to_status_id_str"`
	InReplyToUserID             int64                  `json:"in_reply_to_user_id"`
	InReplyToUserIdStr          string                 `json:"in_reply_to_user_id_str"`
	IsTranslationEnabled        bool                   `json:"is_translation_enabled"`
	Lang                        string                 `json:"lang"`
	Place                       Place                  `json:"place"`
	QuotedStatusID              int64                  `json:"quoted_status_id"`
	QuotedStatusIdStr           string                 `json:"quoted_status_id_str"`
	QuotedStatus                *Tweet                 `json:"quoted_status"`
	PossiblySensitive           bool                   `json:"possibly_sensitive"`
	PossiblySensitiveAppealable bool                   `json:"possibly_sensitive_appealable"`
	RetweetCount                int                    `json:"retweet_count"`
	Retweeted                   bool                   `json:"retweeted"`
	RetweetedStatus             *Tweet                 `json:"retweeted_status"`
	Source                      string                 `json:"source"`
	Scopes                      map[string]interface{} `json:"scopes"`
	Text                        string                 `json:"text"`
	User                        User                   `json:"user"`
	WithheldCopyright           bool                   `json:"withheld_copyright"`
	WithheldInCountries         []string               `json:"withheld_in_countries"`
	WithheldScope               string                 `json:"withheld_scope"`

	//Geo is deprecated
	//Geo                  interface{} `json:"geo"`
}

// CreatedAtTime is a convenience wrapper that returns the Created_at time, parsed as a time.Time struct
func (t Tweet) CreatedAtTime() (time.Time, error) {
	return time.Parse(time.RubyDate, t.CreatedAt)
}

// It may be worth placing these in an additional source file(s)

// Could also use User, since the fields match, but only these fields are possible in Contributor
type Contributor struct {
	Id         int64  `json:"id"`
	IdStr      string `json:"id_str"`
	ScreenName string `json:"screen_name"`
}

type Coordinates struct {
	Coordinates [2]float64 `json:"coordinates"` // Coordinate always has to have exactly 2 values
	Type        string     `json:"type"`
}

type ExtendedTweet struct {
	FullText         string   `json:"full_text"`
	DisplayTextRange []int    `json:"display_text_range"`
	Entities         Entities `json:"entities"`
	ExtendedEntities Entities `json:"extended_entities"`
}

// HasCoordinates is a helper function to easily determine if a Tweet has coordinates associated with it
func (t Tweet) HasCoordinates() bool {
	if t.Coordinates != nil {
		if t.Coordinates.Type == "Point" {
			return true
		}
	}
	return false
}

// The following provide convenience and eliviate confusion about the order of coordinates in the Tweet

// Latitude is a convenience wrapper that returns the latitude easily
func (t Tweet) Latitude() (float64, error) {
	if t.HasCoordinates() {
		return t.Coordinates.Coordinates[1], nil
	}
	return 0, fmt.Errorf("No Coordinates in this Tweet")
}

// Longitude is a convenience wrapper that returns the longitude easily
func (t Tweet) Longitude() (float64, error) {
	if t.HasCoordinates() {
		return t.Coordinates.Coordinates[0], nil
	}
	return 0, fmt.Errorf("No Coordinates in this Tweet")
}

// X is a convenience wrapper which returns the X (Longitude) coordinate easily
func (t Tweet) X() (float64, error) {
	return t.Longitude()
}

// Y is a convenience wrapper which return the Y (Lattitude) corrdinate easily
func (t Tweet) Y() (float64, error) {
	return t.Latitude()
}

func (t *Tweet) extractExtendedTweet() {
	// if the TruncatedText is set, the API does not return an extended tweet
	// we need to manually set the Text in this case
	if len(t.Text) > 0 && len(t.FullText) == 0 {
		t.FullText = t.Text
	}

	if len(t.ExtendedTweet.FullText) > 0 {
		t.DisplayTextRange = t.ExtendedTweet.DisplayTextRange
		t.Entities = t.ExtendedTweet.Entities
		t.ExtendedEntities = t.ExtendedTweet.ExtendedEntities
		t.FullText = t.ExtendedTweet.FullText
	}

	// if the API supplied us with information how to extract the shortened
	// text, extract it
	if len(t.Text) == 0 && len(t.DisplayTextRange) == 2 {
		t.Text = t.FullText[t.DisplayTextRange[0]:t.DisplayTextRange[1]]
	}
	// if the truncated text is still empty then full & truncated text are equal
	if len(t.Text) == 0 {
		t.Text = t.FullText
	}
}

func (t *Tweet) UnmarshalJSON(data []byte) error {
	type Alias Tweet
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t.extractExtendedTweet()
	return nil
}
