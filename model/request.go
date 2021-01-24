package model

//GeocoderRequest - Struct for request
type GeocoderRequest struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	County     string `json:"county"`
	State      string `json:"state"`
	Coutry     string `json:"coutry"`
	PostalCode string `json:"postalCode"`
	MaxResult  int    `json:"maxResult"`
}
