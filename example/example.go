package main

import (
	"fmt"

	geocoder ".."
	model "../model"
)

func main() {
	req := model.GeocoderRequest{
		Address: model.GeocoderAddress{
			City: "Berlin",
		},
		Config: model.GeocoderConfig{
			MaxResult: 0,
			Timeout: 1,
		},
	}

	geo := geocoder.New()
	geo.Timeout = 5
	ret := geo.Search(req)
	fmt.Println(ret)

}
