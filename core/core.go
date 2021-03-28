package core

import (
	"encoding/json"
	"strings"

	api "github.com/darlandieterich/go-geocoder/api"
	model "github.com/darlandieterich/go-geocoder/model"
)

type Object model.Object

type Config model.GeocoderConfig

//Search - Search to get the coordenates
func (c *Config) Search(r model.GeocoderRequestSearch) *Object {
	castObject := (*model.GeocoderConfig)(c)
	req := api.Object{}
	req.Config = *castObject
	returnCast := (*Object)(req.Search(r))
	return returnCast
}

//Reverse - Reverse Search by coordinates
func (c *Config) Reverse(r model.GeocoderRequestReverse) *Object {
	castObject := (*model.GeocoderConfig)(c)
	req := api.Object{}
	req.Config = *castObject
	returnCast := (*Object)(req.Reverse(r))
	return returnCast
}

//Find - Search by element with name and contains value
func (o *Object) Find(element string, value string) *Object {
	if o.Error == nil {
		var matches []map[string]interface{}
		for _, item := range o.Object {
			if item[element] != nil {
				if strings.Contains(strings.ToLower(item[element].(string)), strings.ToLower(value)) {
					matches = append(matches, item)
				}
			}
		}
		o.Object = matches
	}
	return o
}

//ToObject - Mapping the JSON string to interface{}
func (o *Object) ToObject() *Object {
	if o.Error == nil {
		var results []map[string]interface{}
		json.Unmarshal([]byte(o.Result.(string)), &results)
		o.Object = results
	}
	return o
}

func (o *Object) First() *Object {
	if o.Error == nil {
		size := len(o.Object)
		if size > 0 {
			var matches []map[string]interface{}
			o.Object = append(matches, o.Object[0])
		}
	}
	return o
}
