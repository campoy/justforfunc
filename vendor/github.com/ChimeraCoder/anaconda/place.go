package anaconda

type Place struct {
	Attributes  map[string]string `json:"attributes"`
	BoundingBox struct {
		Coordinates [][][]float64 `json:"coordinates"`
		Type        string        `json:"type"`
	} `json:"bounding_box"`
	ContainedWithin []struct {
		Attributes  map[string]string `json:"attributes"`
		BoundingBox struct {
			Coordinates [][][]float64 `json:"coordinates"`
			Type        string        `json:"type"`
		} `json:"bounding_box"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
		FullName    string `json:"full_name"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		PlaceType   string `json:"place_type"`
		URL         string `json:"url"`
	} `json:"contained_within"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	FullName    string `json:"full_name"`
	Geometry    struct {
		Coordinates [][][]float64 `json:"coordinates"`
		Type        string        `json:"type"`
	} `json:"geometry"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	PlaceType string   `json:"place_type"`
	Polylines []string `json:"polylines"`
	URL       string   `json:"url"`
}
