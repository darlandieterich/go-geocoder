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
		City:      "Ijui",
		Country:   "Brasil",
		MaxResult: 1,
	}

	/* reverse := model.GeocoderRequestReverse{
		Lat:  51.05,
		Lon:  13.74,
		Zoom: model.ReverseZoomCountry,
	} */
	g := geocoder.New()
	g.Timeout = 8 //seconds
	ret := g.Search(search).Find("display_name", "Ijui")
	fmt.Println("Result", ret.Result)
	//ret := g.Reverse(reverse).ToObject() //.Find("display_name", "Dresden")
	//fmt.Println("Result", reflect.TypeOf(ret.Object))
	/* for i := range ret.Result {
		fmt.Println(fmt.Sprintf("Result: %d, %s", i, ret.Result[i]["display_name"]))
	} */
	//fmt.Println("Result", ret.Result)
	//fmt.Println("Error", ret.Error)
}
