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

		AwaitingSignature bool `json:"awaiting_signature"`
		Enacted           bool `json:"enacted"`
		// "enacted_at": "2014-12-13",

		HousePassageResult string `json:"house_passage_result"`
		// "house_passage_result_at": "2014-12-12T20:04:00Z",

		SenatePassageResult string `json:"senate_passage_result"`
		// "senate_passage_result_at": "2014-12-13",

		Vetoed       bool   `json:"vetoed"`
		IntroducedOn string `json:"introduced_on"`
		LastActionAt string `json:"last_action_at"`
	} `json:"history"`

	/*
	       "last_version": {
	           "version_code": "enr",
	           "issued_on": "2014-12-16",
	           "version_name": "Enrolled Bill",
	           "bill_version_id": "hjres131-113-enr",
	           "urls": {
	               "html": "http://www.gpo.gov/fdsys/pkg/BILLS-113hjres131enr/html/BILLS-113hjres131enr.htm",
	               "pdf": "http://www.gpo.gov/fdsys/pkg/BILLS-113hjres131enr/pdf/BILLS-113hjres131enr.pdf",
	               "xml": "http://www.gpo.gov/fdsys/pkg/BILLS-113hjres131enr/xml/BILLS-113hjres131enr.xml"
	           },
	           "pages": 1
	       },
	       "last_version_on": "2014-12-16",
	       "last_vote_at": "2014-12-13",
	       "number": 131,
	       "official_title": "Making further continuing appropriations for fiscal year 2015, and for other purposes.",
	       "popular_title": null,
	       "related_bill_ids": [
	           "hjres130-113"
	       ],
	       "short_title": null,
	       "sponsor": {
	           "first_name": "Harold",
	           "last_name": "Rogers",
	           "middle_name": null,
	           "name_suffix": null,
	           "nickname": "Hal",
	           "title": "Rep"
	       },
	       "sponsor_id": "R000395",
	       "urls": {
	           "congress": "http://beta.congress.gov/bill/113th/house-joint-resolution/131",
	           "govtrack": "https://www.govtrack.us/congress/bills/113/hjres131",
	           "opencongress": "http://www.opencongress.org/bill/hjres131-113"
	       },
	       "withdrawn_cosponsors_count": 0
	   },
	*/
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
