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

func GetLegislator(bigId string) *Legislator {
	l := &Legislator{}
	internal.GetURL(l, openstatesRoot, map[string]string{}, "legislators", bigId)
	return l
}

func GetLegislators(criteria map[string]string) *[]Legislator {
	l := []Legislator{}
	internal.GetURL(&l, openstatesRoot, criteria, "legislators")
	return &l
}
