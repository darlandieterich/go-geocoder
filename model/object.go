package model

type Object struct {
	Result interface{}
	Object []map[string]interface{}
	Config GeocoderConfig
	Error error
}
