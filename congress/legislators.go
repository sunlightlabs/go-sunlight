package congress

import (
	"github.com/sunlightlabs/go-sunlight/internal"
)

/**
 * Congress Legislator object. Issued for members of the Congress
 */
type Legislator struct {
	BioguideId  string `json:"bioguide_id"`
	Birthday    string `json:"birthday"`
	Chamber     string `json:"chamber"`
	ContactForm string `json:"contact_form"`
	CrpId       string `json:"crp_id"`
	District    int    `json:"district"`
	FacebookId  string `json:"facebook_id"`
	Fax         string `json:"fax"`

	FecIds     []string `json:"fec_ids"`
	FirstName  string   `json:"first_name"`
	MiddleName string   `json:"middle_name"`
	LastName   string   `json:"last_name"`
	Gender     string   `json:"gender"`

	GovtrackId string `json:"govtrack_id"`
	ICPSRId    int    `json:"icpsr_id"`
	InOffice   bool   `json:"in_office"`

	NameSuffix string `json:"name_suffix"`
	Nickname   string `json:"nickname"`
	OcEmail    string `json:"oc_email"`
	Office     string `json:"office"`
	Party      string `json:"party"`

	Phone     string `json:"phone"`
	State     string `json:"state"`
	StateName string `json:"state_name"`
	TermEnd   string `json:"term_end"`
	TermStart string `json:"term_start"`
	ThomasId  string `json:"thomas_id"`
	Title     string `json:"title"`

	TwitterId   string `json:"twitter_id"`
	VotesmartId int    `json:"votesmart_id"`
	Website     string `json:"website"`
	YoutubeId   string `json:"youtube_id"`
}

type LegislatorResult struct {
	Result
	Results []Legislator `json:"results"`
}

func (l Legislator) GetSponsoredBills() (*BillResult, error) {
	bills, err := BillSearch(map[string]string{
		"sponsor_id": l.BioguideId,
	})
	if err != nil {
		return nil, err
	}
	return bills, nil
}

func GetLegislators(criteria map[string]string) (*LegislatorResult, error) {
	l := LegislatorResult{}
	err := internal.GetURL(&l, congressRoot, criteria, "legislators")
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func GetLegislatorsByLatLon(lat string, lon string, params map[string]string) (*LegislatorResult, error) {
	// was float32, seems to be a mistake.
	// var params = map[string]string{
	// 	"latitude":  fmt.Sprintf("%f", lat),
	// 	"longitude": fmt.Sprintf("%f", lon),
	// }

	params["latitude"] = lat
	params["longitude"] = lon

	l := LegislatorResult{}
	err := internal.GetURL(&l, congressRoot, params, "legislators", "locate")
	if err != nil {
		return nil, err
	}
	return &l, nil
}
