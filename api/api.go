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

const urlAPI string = `https://nominatim.openstreetmap.org/search.php?city=new%20york&polygon_geojson=1&format=jsonv2`

//RequestString - Function to return response in string
func RequestString() string {
	res, _ := http.Get(urlAPI)
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

//RequestObject - Function to return object in time execution
func RequestObject() []map[string]interface{} {
	res := RequestString()
	var results []map[string]interface{}
	json.Unmarshal([]byte(res), &results)
	return results
}

//FormatParameters - Process string parameter
func FormatParameters(r model.GeocoderRequest) string {
	params := make(map[string]string)

	params["accept-language"] = r.Language
	params["street"] = r.Street
	params["city"] = r.City
	params["county"] = r.County
	params["state"] = r.State
	params["country"] = r.Country
	params["postalcode"] = r.PostalCode
	params["maxresult"] = func() string {
		if r.MaxResult == 0 {
			return ""
		} else {
			return strconv.Itoa(r.MaxResult)
		}
	}()

	strParams := fmt.Sprintf("%s", urlAPI)

	for keyP, valueP := range params {
		if strings.TrimSpace(valueP) != "" {
			strParams = fmt.Sprintf("%s&%s=%s", strParams, keyP, valueP)
		}
	}

	return strParams
}
