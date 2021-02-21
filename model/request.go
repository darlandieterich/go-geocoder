package model

//GeocoderRequest - Struct for request
type GeocoderRequest struct {
	Street     string
	City       string
	County     string
	State      string
	Country    string
	PostalCode string
}

//GeocoderConfig - Model to configuration
type GeocoderConfig struct {
	Language  string
	MaxResult uint
	Timeout   uint //seconds
}
