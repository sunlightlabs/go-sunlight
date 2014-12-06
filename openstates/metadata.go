package openstates

import (
	"../internal"
)

type Term struct {
	Name string `json:"name"`
}

/**
 * Open States Legislator Role object. Issued for each role they take on
 * within the Legislative body.
 */
type Metadata struct {
	Abbreviation    string `json:"type"`
	CapitolTimezone string `json:"capitol_timezone"`
	Id              string `json:"id"`
	LegislatureName string `json:"legislature_name"`
	LegislatureUrl  string `json:"legislature_url"`
	Name            string `json:"name"`
	Terms           []Term `json:"terms"`
}

/**
 */
func GetMetadata(state string) (*Metadata, error) {
	t := &Metadata{}
	err := internal.GetURL(t, openstatesRoot, map[string]string{}, "metadata", state)
	if err != nil {
		return nil, err
	}
	return t, nil
}

/**
 */
func GetMetadataList() (*[]Metadata, error) {
	t := &[]Metadata{}
	err := internal.GetURL(t, openstatesRoot, map[string]string{}, "metadata")
	if err != nil {
		return nil, err
	}
	return t, nil
}
