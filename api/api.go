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

func (o *Object) rootRequest(url string) *Object {
	fmt.Println(o.Config.Timeout, url)
	client := http.Client{
		Timeout: time.Duration(o.Config.Timeout) * time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		//fmt.Println(err)
		o.Error = err
		return o
	}
	body, _ := ioutil.ReadAll(res.Body)
	o.Internal = string(body)
	return o
}

//Request - Function to return response json
func (o *Object) Search(r model.GeocoderRequestSearch) *Object {
	objHelper := helper.Object{}
	newURL := objHelper.FormatParametersSearch(r)
	objApi := Object{}
	return objApi.rootRequest(newURL)
}

//Reverse - Function to return response json
func (o *Object) Reverse(r model.GeocoderRequestReverse) *Object {
	objHelper := helper.Object{}
	newURL := objHelper.FormatParametersReverse(r)
	objApi := Object{}
	return objApi.rootRequest(newURL)
}
