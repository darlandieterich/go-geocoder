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
			MaxResult: 1,
			Timeout: 3,
		},
	}
}

//Search - Search to get the coordenates
func (c Config) Search(r model.GeocoderRequest) string {
	r.Config.Language = c.Language
	r.Config.MaxResult = c.MaxResult
	r.Config.Timeout = c.Timeout
	return api.RequestString(r)
}

//ReverseSearch - Reverse Search
func (c Config) ReverseSearch() {

}
