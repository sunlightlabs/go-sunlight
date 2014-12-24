package openstates

import (
	"github.com/sunlightlabs/go-sunlight/internal"
)

type CommitteeMember struct {
	LegId string `json:"leg_id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

type Committee struct {
	Timestamps
	Sources

	AllIds       []string          `json:"all_ids"`
	Chamber      string            `json:"chamber"`
	Committee    string            `json:"committee"`
	Id           string            `json:"id"`
	ParentId     string            `json:"parent_id"`
	State        string            `json:"state"`
	Subcommittee string            `json:"subcommittee"`
	Members      []CommitteeMember `json:"members"`
}

func GetCommittee(bigId string) (*Committee, error) {
	t := &Committee{}
	err := internal.GetURL(t, openstatesRoot, map[string]string{}, "committees", bigId)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func GetCommittees(criteria map[string]string) (*[]Committee, error) {
	t := &[]Committee{}
	err := internal.GetURL(t, openstatesRoot, criteria, "committees")
	if err != nil {
		return nil, err
	}
	return t, nil
}
