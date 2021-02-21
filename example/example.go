package main

import (
	"fmt"

	geocoder ".."
	model "../model"
)

func main() {
	req := model.GeocoderRequest{
			City: "Berlin",
	}
	g := geocoder.New()
	ret := g.Search(req).ToObject()
	fmt.Println(ret.Value)
}
