package anaconda

import "net/url"

type GeoSearchResult struct {
	Result struct {
		Places []struct {
			ID              string `json:"id"`
			URL             string `json:"url"`
			PlaceType       string `json:"place_type"`
			Name            string `json:"name"`
			FullName        string `json:"full_name"`
			CountryCode     string `json:"country_code"`
			Country         string `json:"country"`
			ContainedWithin []struct {
				ID          string    `json:"id"`
				URL         string    `json:"url"`
				PlaceType   string    `json:"place_type"`
				Name        string    `json:"name"`
				FullName    string    `json:"full_name"`
				CountryCode string    `json:"country_code"`
				Country     string    `json:"country"`
				Centroid    []float64 `json:"centroid"`
				BoundingBox struct {
					Type        string        `json:"type"`
					Coordinates [][][]float64 `json:"coordinates"`
				} `json:"bounding_box"`
				Attributes struct {
				} `json:"attributes"`
			} `json:"contained_within"`
			Centroid    []float64 `json:"centroid"`
			BoundingBox struct {
				Type        string        `json:"type"`
				Coordinates [][][]float64 `json:"coordinates"`
			} `json:"bounding_box"`
			Attributes struct {
			} `json:"attributes"`
		} `json:"places"`
	} `json:"result"`
	Query struct {
		URL    string `json:"url"`
		Type   string `json:"type"`
		Params struct {
			Accuracy     float64 `json:"accuracy"`
			Granularity  string  `json:"granularity"`
			Query        string  `json:"query"`
			Autocomplete bool    `json:"autocomplete"`
			TrimPlace    bool    `json:"trim_place"`
		} `json:"params"`
	} `json:"query"`
}

func (a TwitterApi) GeoSearch(v url.Values) (c GeoSearchResult, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/geo/search.json", v, &c, _GET, response_ch}
	return c, (<-response_ch).err
}
