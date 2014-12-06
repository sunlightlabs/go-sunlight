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

/**
 * Get the default Sunlight API Key
 */
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

/**
 * Get the current Sunlight API key
 */
func GetAPIKey() string {
	return apiKey
}

/**
 * Set the current Sunlight API key
 */
func SetAPIKey(key string) {
	apiKey = key
}
