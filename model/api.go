package model

const UrlAPISearch string = `https://nominatim.openstreetmap.org/search.php?polygon_geojson=1&format=jsonv2`

const UrlAPIReverse string = `https://nominatim.openstreetmap.org/reverse.php?&format=jsonv2`

//GeocoderConfig - Model to configuration
type GeocoderConfig struct {
	Language  string
	MaxResult uint
	Timeout   uint //seconds
}