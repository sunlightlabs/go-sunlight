package openstates

import (
	"../internal"
)

var openstatesRoot string = "http://openstates.org/api/v1"

type Source struct {
	Url string
}

type Legislator struct {
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	FullName   string   `json:"full_name"`
	Sources    []Source `json:"sources"`
	Id         string   `json:"id"`
	MiddleName string   `json:"middle_name"`
	State      string   `json:"state"`
	Active     bool     `json:"active"`
	PhotoUrl   string   `json:"photo_url"`
	Url        string   `json:"url"`
}

func GetLegislator(bigId string) (*Legislator, error) {
	l := &Legislator{}
	err := internal.GetURL(l, openstatesRoot, map[string]string{}, "legislators", bigId)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func GetLegislators(criteria map[string]string) (*[]Legislator, error) {
	l := []Legislator{}
	err := internal.GetURL(&l, openstatesRoot, criteria, "legislators")
	if err != nil {
		return nil, err
	}
	return &l, nil
}
