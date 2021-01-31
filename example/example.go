package main

import (
	"fmt"

	api "../api"
	model "../model"
)

func main() {
	//res := api.RequestObject()
	//fmt.Println(res[0]["licence"])
	req := model.GeocoderRequest{}
	req.City = "Berlin"
	req.MaxResult = 5
	ret := api.FormatParameters(req)
	fmt.Println(ret)

}
