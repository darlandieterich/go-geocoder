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
	g.Timeout = 3 //seconds
	ret := g.Search(req).ToObject()
	fmt.Println(ret.Value)
	fmt.Println(ret.Error)
}
