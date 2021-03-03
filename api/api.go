package api

import (
	"fmt"
	"time"
	"io/ioutil"
	"net/http"

	model "../model"
	helper "../helper"
)

type Object model.Object

//Request - Function to return response json
func (o *Object) Request(r model.GeocoderRequestSearch) *Object {
	objHelper := helper.Object{}
	newURL := objHelper.FormatParametersSearch(r)
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
