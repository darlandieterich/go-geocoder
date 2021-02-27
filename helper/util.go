package helper

import (
	"strconv"
	"strings"
	"fmt"
	model "../model"
)

type Object model.Object

//FormatParameters - Process string parameter
func (o *Object) FormatParameters(r model.GeocoderRequest) string {
	params := make(map[string]string)

	params["street"] = r.Street
	params["city"] = r.City
	params["county"] = r.County
	params["state"] = r.State
	params["country"] = r.Country
	params["postalcode"] = r.PostalCode
	params["accept-language"] = o.Config.Language
	if o.Config.MaxResult > 0 {
		params["maxresult"] = strconv.Itoa(int(o.Config.MaxResult))
	}

	strParams := fmt.Sprintf("%s", model.UrlAPI)

	for keyP, valueP := range params {
		if strings.TrimSpace(valueP) != "" {
			strParams = fmt.Sprintf("%s&%s=%s", strParams, keyP, valueP)
		}
	}

	return strParams
}
