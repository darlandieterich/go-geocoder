package model

type Object struct {
	Internal string
	Result []map[string]interface{}
	Config GeocoderConfig
	Error error
}
