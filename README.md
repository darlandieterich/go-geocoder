# go-geocoder
Client for geographic location. Define the location and return the geographic coordinates or reverse.

## Example
```go
	req := model.GeocoderRequest{
			City: "Berlin",
	}
	g := geocoder.New()
	g.Timeout = 3 //seconds
	ret := g.Search(req).ToObject()
	fmt.Println(ret.Value)
	fmt.Println(ret.Error)
```