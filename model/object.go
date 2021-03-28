package model

type Object struct {
	Object []map[string]interface{}
	Config GeocoderConfig
	Error  error
}
