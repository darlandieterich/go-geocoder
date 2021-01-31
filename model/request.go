package model

//GeocoderRequest - Struct for request
type GeocoderRequest struct {
	Language   string `json:"accept-language" default0:"en"`
	Street     string `json:"street"`
	City       string `json:"city"`
	County     string `json:"county"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
	MaxResult  int    `json:"maxResult"`
}
