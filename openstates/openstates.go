package openstates

import (
	"../internal"
)

// Open States API root URL.
var openstatesRoot string = "http://openstates.org/api/v1"

/**
 * Open States data Source object. Present on nearly all objects in the
 * Open States database.
 */
type Source struct {
	Url string
}

/**
 * Open States Legislator object. Issued for members of the Legislature.
 */
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

/**
 * Get an Open States Legislator by their Open States issued ID, such as
 * MAL000155
 */
func GetLegislator(bigId string) (*Legislator, error) {
	l := &Legislator{}
	err := internal.GetURL(l, openstatesRoot, map[string]string{}, "legislators", bigId)
	if err != nil {
		return nil, err
	}
	return l, nil
}

/**
 * Query the Open States API by params defined in
 * http://sunlightlabs.github.io/openstates-api/legislators.html#methods/legislator-search
 */
func GetLegislators(criteria map[string]string) (*[]Legislator, error) {
	l := []Legislator{}
	err := internal.GetURL(&l, openstatesRoot, criteria, "legislators")
	if err != nil {
		return nil, err
	}
	return &l, nil
}
