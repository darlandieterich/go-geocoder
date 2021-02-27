package model

type Object struct {
	Internal interface{}
	Result []map[string]interface{}
	Config GeocoderConfig
	Error error
}
