package congress

import (
	"fmt"
	"github.com/sunlightlabs/go-sunlight/internal"
)

type District struct {
	State    string `json:"state"`
	District int    `json:"district"`
}

type DistrictResult struct {
	Result
	Results []District `json:"results"`
}

func GetDistrictsByLatLon(lat float32, lon float32) (*DistrictResult, error) {
	var params = map[string]string{
		"latitude":  fmt.Sprintf("%f", lat),
		"longitude": fmt.Sprintf("%f", lon),
	}

	l := DistrictResult{}
	err := internal.GetURL(&l, congressRoot, params, "districts", "locate")
	if err != nil {
		return nil, err
	}
	return &l, nil
}
