package api

//example - https://nominatim.openstreetmap.org/search.php

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	model "../model"
)

const urlAPI string = `https://nominatim.openstreetmap.org/search.php?polygon_geojson=1&format=jsonv2`

//RequestString - Function to return response in string
func RequestString(r model.GeocoderRequest) string {
	newURL := FormatParameters(r)
	res, _ := http.Get(newURL)
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

//RequestObject - Function to return object in time execution
func RequestObject(r model.GeocoderRequest) []map[string]interface{} {
	res := RequestString(r)
	var results []map[string]interface{}
	json.Unmarshal([]byte(res), &results)
	return results
}

//FormatParameters - Process string parameter
func FormatParameters(r model.GeocoderRequest) string {
	params := make(map[string]string)

	params["street"] = r.Address.Street
	params["city"] = r.Address.City
	params["county"] = r.Address.County
	params["state"] = r.Address.State
	params["country"] = r.Address.Country
	params["postalcode"] = r.Address.PostalCode
	params["accept-language"] = r.Config.Language
	if r.Config.MaxResult > 0 {
		params["maxresult"] = strconv.Itoa(int(r.Config.MaxResult))
	}

	strParams := fmt.Sprintf("%s", urlAPI)

	for keyP, valueP := range params {
		if strings.TrimSpace(valueP) != "" {
			strParams = fmt.Sprintf("%s&%s=%s", strParams, keyP, valueP)
		}
	}

	return strParams
}
