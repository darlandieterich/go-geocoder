package model


type Object struct {
	internal interface{}
	Result []map[string]interface{}
	Config GeocoderConfig
	Error error
}