package geocoder

import (
	"fmt"

	"./api"
	"./model"
)

type Config struct {
	cnf model.GeocoderConfig
}

//New - declaration
func New(n model.GeocoderAddress) {
	fmt.Println("oe")
}

//Search - Search to get the coordenates
func (c Config) Search(r model.GeocoderRequest) string {
	r.Config = c.cnf
	return api.RequestString(r)
}

//ReverseSearch - Reverse Search
func (c Config) ReverseSearch() {

}
