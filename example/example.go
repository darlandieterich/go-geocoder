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
			MaxResult: 2,
		},
	}

	geo := geocoder.New()
	geo.MaxResult = 5
	ret := geo.Search(req)
	fmt.Println(ret)

}
