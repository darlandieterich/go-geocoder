package model

//GeocoderRequestReverse - Struct for reverse
type GeocoderRequestReverse struct {
	Lat  float32
	Lon  float32
	Zoom uint //Result limitation - 0-18

	/* zoom	address detail
	3	country
	5	state
	8	county
	10	city
	14	suburb
	16	major streets
	17	major and minor streets
	18	building */
}