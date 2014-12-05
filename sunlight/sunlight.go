package sunlight

var (
	apiKey string
)

func GetAPIKey() string {
	return apiKey
}

func SetAPIKey(key string) {
	apiKey = key
}
