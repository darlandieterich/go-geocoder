package main

import (

	//"net/http"
	//"io/ioutil"
	//"encoding/json"
	//"time"

	"fmt"

	geocoder "github.com/darlandieterich/go-geocoder"
	model "github.com/darlandieterich/go-geocoder/model"
)

func main() {
	search := model.GeocoderRequestSearch{
		City:      "Dresden",
		Country:   "Germany",
		MaxResult: 1,
	}

	reverse := model.GeocoderRequestReverse{
		Lat:  51.05,
		Lon:  13.74,
		Zoom: model.ReverseZoomCountry,
	}
	g := geocoder.New()
	g.Timeout = 8           //seconds
	ret := g.Search(search) //.Find("display_name", "Dresden")
	fmt.Println("Result", ret.Result)
	ret = g.Reverse(reverse) //.Find("display_name", "Dresden")
	fmt.Println("Result", ret.Result)
	/* for i := range ret.Result {
		fmt.Println(fmt.Sprintf("Result: %d, %s", i, ret.Result[i]["display_name"]))
	} */
	//fmt.Println("Result", ret.Result)
	//fmt.Println("Error", ret.Error)
}
