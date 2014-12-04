package internal

import (
	"net/url"
	"strings"
)

func GenerateURL(root string, resource ...string) string {
	ret := root + "/" + strings.Join(resource, "/")
	return ret
}

func GetAPIKey() string {
	return "foo"
}

func SetAPIKey() {

}

func QueryURI(root string, params map[string]string, resource ...string) string {
	apikey := GetAPIKey()
	uri := GenerateURL(root, resource...) + "?apikey=" + apikey
	for k, v := range params {
		uri = uri + "&" + url.QueryEscape(k) + "=" + url.QueryEscape(v)
	}
	return uri
}
