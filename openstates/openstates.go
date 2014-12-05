package openstates

import (
	"../internal"
)

var openstatesRoot string = "http://openstates.org/api/v1/"

type Legislator struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func GetLegislator(bigId string) *Legislator {
	l := &Legislator{}
	internal.GetURL(l, openstatesRoot, map[string]string{}, "legislators", bigId)
	return l
}
