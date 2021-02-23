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
	g.Timeout = 8 //seconds
	ret := g.Search(req).ToObject().Find("display_name", "Ijui")
	/* m := ret.Result
	fmt.Println(m) */
	/* for i := range m {
		fmt.Println(fmt.Sprintf("Result: %d, %s", i, m[i]["display_name"]))
	} */
	fmt.Println(ret.Error)
}
