package congress

import (
	"github.com/sunlightlabs/go-sunlight/internal"
)

type Bill struct {
	BillId          string   `json:"bill_id"`
	BillType        string   `json:"bill_type"`
	Chamber         string   `json:"chamber"`
	CommitteeIds    []string `json:"committee_ids"`
	Congress        int      `json:"congress"`
	CosponsorsCount int      `json:"cosponsors_count"`
	EnactedAs       struct {
		Congress int    `json:"congress"`
		LawType  string `json:"law_type"`
		Number   int    `json:"number"`
	} `json:"enacted_as"`
	History struct {
		Active bool `json:"active"`
		// "active_at": "2014-12-12T20:04:00Z",

		AwaitingSignature bool   `json:"awaiting_signature"`
		Enacted           bool   `json:"enacted"`
		EnactedAt         string `json:"enacted_at"`

		HousePassageResult string `json:"house_passage_result"`
		// "house_passage_result_at": "2014-12-12T20:04:00Z",

		SenatePassageResult string `json:"senate_passage_result"`
		// "senate_passage_result_at": "2014-12-13",

		Vetoed bool `json:"vetoed"`
	} `json:"history"`

	LastVersion struct {
		VersionCode   string            `json:"version_code"`
		IssuedOn      string            `json:"issued_on"`
		VersionName   string            `json:"version_name"`
		BillVersionId string            `json:"bill_version_id"`
		Urls          map[string]string `json:"urls"`
		Pages         int               `json:"pages"`
	} `json:"last_version"`

	LastVersionOn string `json:"last_version_on"`
	LastVoteAt    string `json:"last_vote_at"`
	IntroducedOn  string `json:"introduced_on"`
	LastActionAt  string `json:"last_action_at"`

	Number         int      `json:"number"`
	OfficialTitle  string   `json:"official_title"`
	PopularTitle   string   `json:"popular_title"`
	RelatedBillIds []string `json:"related_bill_ids"`
	ShortTitle     string   `json:"short_title"`
	Sponsor        struct {
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		MiddleName string `json:"middle_name"`
		NameSuffix string `json:"name_suffix"`
		Nickname   string `json:"nickname"`
		Title      string `json:"title"`
	} `json:"sponsor"`
	SponsorId                string            `json:"sponsor_id"`
	Urls                     map[string]string `json:"urls"`
	WithdrawnCosponsorsCount int               `json:"withdrawn_cosponsors_count"`
}

type BillResult struct {
	Result
	Results []Bill `json:"results"`
}

func BillSearch(criteria map[string]string) (*BillResult, error) {
	l := BillResult{}
	err := internal.GetURL(&l, congressRoot, criteria, "bills")
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func BillTextSearch(query string, criteria map[string]string) (*BillResult, error) {
	criteria["query"] = query
	l := BillResult{}
	err := internal.GetURL(&l, congressRoot, criteria, "bills", "search")
	if err != nil {
		return nil, err
	}
	return &l, nil
}
