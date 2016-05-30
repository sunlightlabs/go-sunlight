package congress

import (
	"github.com/sunlightlabs/go-sunlight/internal"
)

type Vote struct {
	BillId       string `json:"bill_id"`
	AmendmentId  string `json:"amendment_id"`
	NominationId string `json:"nomination_id"`
	Chamber      string `json:"chamber"`
	Congress     int    `json:"congress"`
	Number       int    `json:"number"`
	Question     string `json:"question"`
	Required     string `json:"required"`
	Result       string `json:"result"`
	RollId       string `json:"roll_id"`
	RollType     string `json:"roll_type"`
	Source       string `json:"source"`
	Url          string `json:"url"`
	VoteType     string `json:"vote_type"`
	VotedAt      string `json:"voted_at"`
	Year         int    `json:"year"`
}

type VoteResult struct {
	Result
	Results []Vote `json:"results"`
}

func GetVotes(criteria map[string]string) (*VoteResult, error) {
	l := VoteResult{}
	err := internal.GetURL(&l, congressRoot, criteria, "votes")
	if err != nil {
		return nil, err
	}
	return &l, nil
}
