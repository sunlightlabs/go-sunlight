package internal

import (
	"../sunlight"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func GenerateURL(root string, resource ...string) string {
	ret := root + "/" + strings.Join(resource, "/")
	return ret
}

func QueryURL(root string, params map[string]string, resource ...string) string {
	apikey := sunlight.GetAPIKey()
	uri := GenerateURL(root, resource...) + "?apikey=" + apikey
	for k, v := range params {
		uri = uri + "&" + url.QueryEscape(k) + "=" + url.QueryEscape(v)
	}
	return uri
}

func GetURL(target interface{}, root string, params map[string]string, resource ...string) {
	url := QueryURL(root, params, resource...)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, target)
}
