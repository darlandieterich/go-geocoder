package model

//UrlAPISearch - Url to request Search
const UrlAPISearch string = `https://nominatim.openstreetmap.org/search.php?polygon_geojson=1&format=jsonv2`

//UrlAPIReverse - Url to request reverse Search
const UrlAPIReverse string = `https://nominatim.openstreetmap.org/reverse.php?&format=jsonv2`

//GeocoderConfig - Model to configuration
type GeocoderConfig struct {
	Language  string
	Timeout   uint //seconds
}