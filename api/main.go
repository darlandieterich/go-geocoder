package api

//example - https://nominatim.openstreetmap.org/search.php

import (
	"io/ioutil"
	"net/http"
)

func Request() string {
	url := "https://nominatim.openstreetmap.org/search.php?city=new%20york&polygon_geojson=1&format=jsonv2"
	res, _ := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
