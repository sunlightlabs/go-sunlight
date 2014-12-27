package congress

var congressRoot string = "http://congress.api.sunlightfoundation.com"

type Result struct {
	Count int `json:"count"`
	Page  struct {
		Count   int `json:"count"`
		PerPage int `json:"per_page"`
		Page    int `json:"page"`
	} `json:"page"`
}
