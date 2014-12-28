package congress

import (
	"github.com/sunlightlabs/go-sunlight/internal"
)

type Hearing struct {
	CommitteeId    string   `json:"committee_id"`
	Congress       int      `json:"congress"`
	Room           string   `json:"room"`
	Dc             bool     `json:"dc"`
	BillIds        []string `json:"bill_ids"`
	Chamber        string   `json:"chamber"`
	SubcommitteeId string   `json:"subcommittee_id"`
	HearingId      string   `json:"hearing_id"`
	OccursAt       string   `json:"occurs_at"`
	Description    string   `json:"description"`
}

type HearingResult struct {
	Result
	Results []Hearing `json:"results"`
}

func GetHearings(criteria map[string]string) (*HearingResult, error) {
	l := HearingResult{}
	err := internal.GetURL(&l, congressRoot, criteria, "hearings")
	if err != nil {
		return nil, err
	}
	return &l, nil
}
