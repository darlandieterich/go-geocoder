package geocoder

import (
	"./api"
	"./model"
)

type Config model.GeocoderConfig

//New - declaration of default values
func New() Config {
	return Config{
		Language:  "en",
		MaxResult: 1,
		Timeout: 5,
	}
}

//Search - Search to get the coordenates
func (c *Config) Search(r model.GeocoderRequest) *api.ApiObject {
	castObject := (*model.GeocoderConfig)(c)
	req := api.ApiObject{}
	req.Config = *castObject
	return req.Request(r)
}

//ReverseSearch - Reverse Search
func (c *Config) ReverseSearch() {
}
