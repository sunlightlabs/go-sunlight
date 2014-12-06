package openstates

import (
	"../internal"
)

/**
 * Open States Legislator Role object. Issued for each role they take on
 * within the Legislative body.
 */
type Term struct {
	Abbreviation    string `json:"type"`
	CapitolTimezone string `json:"capitol_timezone"`
	Id              string `json:"id"`
	LegislatureName string `json:"legislature_name"`
	LegislatureUrl  string `json:"legislature_url"`
	Name            string `json:"name"`
}

/**
 */
func GetMetadata(state string) (*Term, error) {
	t := &Term{}
	err := internal.GetURL(t, openstatesRoot, map[string]string{}, "metadata", state)
	if err != nil {
		return nil, err
	}
	return t, nil
}
