package sunlight

import (
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

var (
	apiKey string = getDefaultAPIKey()
)

func getDefaultAPIKey() string {
	ekey := os.Getenv("SUNLIGHT_API_KEY")
	if ekey != "" {
		return ekey
	}

	usr, _ := user.Current()
	dir := usr.HomeDir

	fkey, err := ioutil.ReadFile(dir + "/.sunlight.key")

	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(fkey))
}

func GetAPIKey() string {
	return apiKey
}

func SetAPIKey(key string) {
	apiKey = key
}
