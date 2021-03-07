package helper

import (
	"strings"
	"fmt"
	model "../model"
)

type Object model.Object

//FormatParametersSearch - Process string parameter
func (o *Object) FormatParametersSearch(r model.GeocoderRequestSearch) string {
	params := make(map[string]string)

	params["street"] = r.Street
	params["city"] = r.City
	params["county"] = r.County
	params["state"] = r.State
	params["country"] = r.Country
	params["postalcode"] = r.PostalCode
	params["accept-language"] = o.Config.Language
	if r.MaxResult > 0 {
		params["maxresult"] = fmt.Sprint(r.MaxResult)
	}

	strParams := fmt.Sprintf("%s", model.UrlAPISearch)

	return formatParameters(strParams, params)
}

//FormatParametersReverse - Process parameters to return string formated
func (o *Object) FormatParametersReverse(r model.GeocoderRequestReverse) string {
	params := make(map[string]string)

	params["lat"] = fmt.Sprint(r.Lat)
	params["lon"] = fmt.Sprint(r.Lon)
	params["zoom"] = fmt.Sprint(r.Zoom)

	strParams := fmt.Sprintf("%s", model.UrlAPIReverse)

	return formatParameters(strParams, params)
}

//formatParameters - Internal function to join the parameters
func formatParameters(strParams string, params map[string]string) string {
	for keyP, valueP := range params {
		if strings.TrimSpace(valueP) != "" {
			strParams = fmt.Sprintf("%s&%s=%s", strParams, keyP, valueP)
		}
	}
	return strParams
}
