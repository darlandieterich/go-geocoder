# go-geocoder
Client for geographic location. Define the location and return the geographic coordinates or reverse.

## Example
```go
	req := model.GeocoderRequest{
		Address: model.GeocoderAddress{
			City: "Berlin",
		},
		Config: model.GeocoderConfig{
			MaxResult: 0,
			Timeout: 1,
		},
	}

	geo := geocoder.New()
	geo.Timeout = 2
	ret := geo.Search(req)
	fmt.Println(ret)
```