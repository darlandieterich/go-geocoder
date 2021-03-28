package geocoder

import (
	core "github.com/darlandieterich/go-geocoder/core"
	model "github.com/darlandieterich/go-geocoder/model"
)

type Config model.GeocoderConfig

type Object model.Object

//New - declaration of default values
func New() Config {
	return Config{
		Language: "en",
		Timeout:  3,
	}
}

func (c *Config) Search(r model.GeocoderRequestSearch) *Object {
	ca := (*core.Config)(c)
	return (*Object)(ca.Search(r))
}

func (c *Config) Reverse(r model.GeocoderRequestReverse) *Object {
	ca := (*core.Config)(c)
	return (*Object)(ca.Reverse(r))
}

func (o *Object) Find(element string, value string) *Object {
	co := (*core.Object)(o)
	return (*Object)(co.Find(element, value))
}

func (o *Object) First() *Object {
	co := (*core.Object)(o)
	return (*Object)(co.First())
}

func (o *Object) Last() *Object {
	co := (*core.Object)(o)
	return (*Object)(co.Last())
}
