package model

//GeocoderRequest - Main struct of request
type GeocoderRequest struct {
	Address GeocoderAddress
	Config  GeocoderConfig
}

//GeocoderAddress - Struct for request
type GeocoderAddress struct {
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
