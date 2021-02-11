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
		},
	}

	geo := geocoder.New()
	ret := geo.Search(req)
	fmt.Println(ret)

}
