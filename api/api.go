package api

//example - https://nominatim.openstreetmap.org/search.php

import (
	"fmt"
	"time"
	"io/ioutil"
	"net/http"

	model "../model"
	helper "../helper"
)

type Object model.Object

//Request - Function to return response in string
func (o *Object) Request(r model.GeocoderRequest) *Object {
	objHelper := helper.Object{}
	newURL := objHelper.FormatParameters(r)
	fmt.Println(o.Config.Timeout, newURL)
	client := http.Client{
		Timeout: time.Duration(o.Config.Timeout) * time.Second,
	}
	res, err := client.Get(newURL)
	if err != nil {
		//fmt.Println(err)
		o.Error = err
		return o
	}
	body, _ := ioutil.ReadAll(res.Body)
	o.Internal = string(body)
	return o
}
