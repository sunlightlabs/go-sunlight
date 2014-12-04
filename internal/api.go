package internal

import (
	"net/url"
	"strings"
)

var (
	apiKey string
)

func GenerateURL(root string, resource ...string) string {
	ret := root + "/" + strings.Join(resource, "/")
	return ret
}

func QueryURL(root string, params map[string]string, resource ...string) string {
	apikey := GetAPIKey()
	uri := GenerateURL(root, resource...) + "?apikey=" + apikey
	for k, v := range params {
		uri = uri + "&" + url.QueryEscape(k) + "=" + url.QueryEscape(v)
	}
	return uri
}

func GetAPIKey() string {
	return apiKey
}

func SetAPIKey(key string) {
	apiKey = key
}
