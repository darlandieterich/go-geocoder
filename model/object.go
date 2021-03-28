package model

type Object struct {
	Result []map[string]interface{}
	Config GeocoderConfig
	Error  error
}
