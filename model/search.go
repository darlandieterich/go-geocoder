package model

//GeocoderRequest - Struct for request
type GeocoderRequestSearch struct {
	Street     string
	City       string
	County     string
	State      string
	Country    string
	PostalCode string
}

//GeocoderConfig - Model to configuration
type GeocoderConfig struct {
	Language  string
	MaxResult uint
	Timeout   uint //seconds
}

//GeocoderResponse - Stack Struct for response
type GeocoderResponse struct {
	Response []GeocoderResponseObject `json:"[]"`
}

// GeocoderResponseObject - Struct for response
type GeocoderResponseObject struct {
	OsmType string `json:"osm_type"`
	OsmID   string `json:"osm_id"`
}

/* "place_id": 259048015,
"licence": "Data © OpenStreetMap contributors, ODbL 1.0. https://osm.org/copyright",
"osm_type": "relation",
"osm_id": 175905,
"boundingbox": [
	"40.477399",
	"40.9161785",
	"-74.25909",
	"-73.7001809"
],
"lat": "40.7127281",
"lon": "-74.0060152",
"display_name": "Nova Iorque, Estados Unidos da América",
"place_rank": 10,
"category": "boundary",
"type": "administrative",
"importance": 0.8175766114518461,
"icon": "https://nominatim.openstreetmap.org/ui/mapicons//poi_boundary_administrative.p.20.png",
"geojson": {} */
