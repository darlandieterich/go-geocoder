# go-geocoder
Client for geographic location. Define the location and return the geographic coordinates or reverse.

## Example
```go
req := model.GeocoderRequest{
  Address: model.GeocoderAddress{
    City: "Berlin",
  },
  Config: model.GeocoderConfig{
    MaxResult: 2,
  },
}

geo := geocoder.New()
ret := geo.Search(req)
fmt.Println(ret)
```