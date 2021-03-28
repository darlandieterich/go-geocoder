package core

import (
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

//First - Return the first position of collection
func (o *Object) First() *Object {
	if o.Error == nil {
		size := len(o.Result)
		if size > 0 {
			var first []map[string]interface{}
			o.Result = append(first, o.Result[0])
		}
	}
	return o
}

//Last - Return the last position of collection
func (o *Object) Last() *Object {
	if o.Error == nil {
		size := len(o.Result)
		if size > 0 {
			var last []map[string]interface{}
			o.Result = append(last, o.Result[size-1])
		}
	}
	return o
}
