package api

//example - https://nominatim.openstreetmap.org/search.php

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const urlAPI string = `https://nominatim.openstreetmap.org/search.php?city=new%20york&polygon_geojson=1&format=jsonv2`

func RequestString() string {
	res, _ := http.Get(urlAPI)
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func RequestObject() []map[string]interface{} {
	res := RequestString()
	var results []map[string]interface{}
	json.Unmarshal([]byte(res), &results)
	return results
}
