# go-geocoder
Client for geographic location. Define the location and return the geographic coordinates or reverse.
This project base on API openstreetmap.org.

## Example
```go
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
```