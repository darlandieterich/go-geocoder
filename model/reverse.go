package model

//Limitation zoom address detail
const (
	ReverseZoomCountry = 3
	ReverseZoomState = 5
	ReverseZoomCounty = 8
	ReverseZoomCity = 10
	ReverseZoomSuburd = 14
	ReverseZoomMajorStreets = 16
	ReverseZoomMajorMinorStreets = 17
	ReverseZoomBuilding = 18
)

//GeocoderRequestReverse - Struct for reverse
type GeocoderRequestReverse struct {
	Lat  float32
	Lon  float32
	Zoom uint
}