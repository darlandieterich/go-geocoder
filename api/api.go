package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	helper "github.com/darlandieterich/go-geocoder/helper"
	model "github.com/darlandieterich/go-geocoder/model"
)

type Object model.Object

//rootRequest - Internal function to request and return result
func (o *Object) rootRequest(url string, inCollection bool) *Object {
	fmt.Println(o.Config.Timeout, url)
	client := http.Client{
		Timeout: time.Duration(o.Config.Timeout) * time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		o.Error = err
		return o
	}
	body, _ := ioutil.ReadAll(res.Body)

	var ret = ""
	if inCollection {
		ret = "[" + string(body) + "]"
	} else {
		ret = string(body)
	}

	var results []map[string]interface{}
	json.Unmarshal([]byte(ret), &results)

	o.Object = results
	return o
}

//Request - Function to return response json
func (o *Object) Search(r model.GeocoderRequestSearch) *Object {
	objHelper := helper.Object{}
	newURL := objHelper.FormatParametersSearch(r)
	objApi := Object{}
	return objApi.rootRequest(newURL, false)
}

//Reverse - Function to return response json
func (o *Object) Reverse(r model.GeocoderRequestReverse) *Object {
	objHelper := helper.Object{}
	newURL := objHelper.FormatParametersReverse(r)
	objApi := Object{}
	return objApi.rootRequest(newURL, true)
}
