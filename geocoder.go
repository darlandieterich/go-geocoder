package geocoder

import (
	model "./model"
	core "./core"
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

func (c *Config) Search(r model.GeocoderRequestSearch) *Object {
	ca := (*core.Config)(c)
	return (*Object)(ca.Search(r))
}

func (o *Object) ToObject() *Object {
	co := (*core.Object)(o)
	return (*Object)(co.ToObject())
}

func (o *Object) Find(element string, value string) *Object {
	co := (*core.Object)(o)
	return (*Object)(co.Find(element, value))
}