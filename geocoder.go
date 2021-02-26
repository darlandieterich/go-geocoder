package geocoder

import (
	"./api"
	"./model"
  "strings"
	"encoding/json"
)

type Config model.GeocoderConfig

type Object model.Object

//New - declaration of default values
func New() Config {
	return Config{
		Language:  "en",
		MaxResult: 1,
		Timeout: 5,
	}
}

//Search - Search to get the coordenates
func (c *Config) Search(r model.GeocoderRequest) *Object {
	castObject := (*model.GeocoderConfig)(c)
	req := api.Object{}
	req.Config = *castObject
	returnCast := (*Object)(req.Request(r))
	return returnCast
}

//Find - Search by element with name and contains value
func (o *Object) Find(element string, value string) *Object {
	if o.Error == nil {
		var matches []map[string]interface{}
		for _, item := range o.Result {
			if item[element] != nil {
				if strings.Contains(strings.ToLower(item[element].(string)), strings.ToLower(value)) {
					matches = append(matches, item)
				}
			}
		}
		o.Result = matches
	}
	return o
}

//RequestObject - Function to return object in time execution
func (o *Object) ToObject() *Object {
	if o.Error == nil {
		var results []map[string]interface{}
		json.Unmarshal([]byte(o.Internal.(string)), &results)
		o.Result = results
	}
	return o
}

//ReverseSearch - Reverse Search
func (c *Config) ReverseSearch() {

}
