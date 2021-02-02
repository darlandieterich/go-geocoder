package geocoder

import (
	"./api"
	"./model"
)

type Config struct {
	model.GeocoderConfig
}

//New - declaration
func New() Config {
	return Config{
		model.GeocoderConfig{
			Language:  "en",
			MaxResult: 0,
		},
	}
}

//Search - Search to get the coordenates
func (c Config) Search(r model.GeocoderRequest) string {
	r.Config.Language = c.Language
	r.Config.MaxResult = c.MaxResult
	return api.RequestString(r)
}

//ReverseSearch - Reverse Search
func (c Config) ReverseSearch() {

}
