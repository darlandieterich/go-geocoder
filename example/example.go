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

	test := geocoder.Config{}
	ret := test.Search(req)
	fmt.Println(ret)

}
