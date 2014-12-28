package congress

import (
	"github.com/sunlightlabs/go-sunlight/internal"
)

type Committee struct {
	Chamber           string `json:"chamber"`
	CommitteeId       string `json:"committee_id"`
	Name              string `json:"name"`
	ParentCommitteeId string `json:"parent_committee_id"`
	Subcommittee      bool   `json:"subcommittee"`
}

type CommitteeResult struct {
	Result
	Results []Committee `json:"results"`
}

func GetCommittees(criteria map[string]string) (*CommitteeResult, error) {
	l := CommitteeResult{}
	err := internal.GetURL(&l, congressRoot, criteria, "committees")
	if err != nil {
		return nil, err
	}
	return &l, nil
}
