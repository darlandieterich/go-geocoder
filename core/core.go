package core

import (
  "strings"
	"encoding/json"

	api   "../api"
	model "../model"
)

type Object model.Object

type Config model.GeocoderConfig

//Search - Search to get the coordenates
func (c *Config) Search(r model.GeocoderRequestSearch) *Object {
	castObject := (*model.GeocoderConfig)(c)
	req := api.Object{}
	req.Config = *castObject
	returnCast := (*Object)(req.Request(r))
	return returnCast
}

//ReverseSearch - Reverse Search
func (c *Config) ReverseSearch() *Object {
 return &Object{}
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

//ToObject - Mapping the JSON string to interface{}
func (o *Object) ToObject() *Object {
	if o.Error == nil {
		var results []map[string]interface{}
		json.Unmarshal([]byte(o.Internal.(string)), &results)
		o.Result = results
	}
	return o
}