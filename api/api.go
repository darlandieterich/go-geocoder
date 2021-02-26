package api

//example - https://nominatim.openstreetmap.org/search.php

import (
	/* "encoding/json" */
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	model "../model"
)

type Object model.Object

const urlAPI string = `https://nominatim.openstreetmap.org/search.php?polygon_geojson=1&format=jsonv2`

//Request - Function to return response in string
func (o *Object) Request(r model.GeocoderRequest) *Object {
	newURL := o.FormatParameters(r)
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

//FormatParameters - Process string parameter
func (o *Object) FormatParameters(r model.GeocoderRequest) string {
	params := make(map[string]string)

	params["street"] = r.Street
	params["city"] = r.City
	params["county"] = r.County
	params["state"] = r.State
	params["country"] = r.Country
	params["postalcode"] = r.PostalCode
	params["accept-language"] = o.Config.Language
	if o.Config.MaxResult > 0 {
		params["maxresult"] = strconv.Itoa(int(o.Config.MaxResult))
	}

	strParams := fmt.Sprintf("%s", urlAPI)

	for keyP, valueP := range params {
		if strings.TrimSpace(valueP) != "" {
			strParams = fmt.Sprintf("%s&%s=%s", strParams, keyP, valueP)
		}
	}

	return strParams
}
