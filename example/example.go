package main

import (
	"fmt"

	geocoder ".."
	model "../model"
)

func main() {
	req := model.GeocoderRequest{
			City: "Panambi",
			Country: "Brasil",
	}
	g := geocoder.New()
	g.Timeout = 5 //seconds
	ret := g.Search(req).ToObject()
	m := ret.Result
	for i := range m {
		fmt.Println(fmt.Sprintf("Result: %d, %s", i, m[i]["display_name"]))
	}
	fmt.Println(ret.Error)
}
