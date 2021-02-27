package main

import (
	"fmt"

	geocoder ".."
	model "../model"
)

func main() {
	req := model.GeocoderRequest{
			City: "Ijui",
			Country: "Brasil",
	}
	g := geocoder.New()
	g.Timeout = 8 //seconds
	ret := g.Search(req).ToObject() //.Find("display_name", "Ijui")

	for i := range ret.Result {
		fmt.Println(fmt.Sprintf("Result: %d, %s", i, ret.Result[i]["display_name"]))
	}
	//fmt.Println("Result", ret.Result)
	fmt.Println("Error", ret.Error)
}
