package congress

import (
	"github.com/sunlightlabs/go-sunlight/internal"
)

type UpcomingBill struct {
	BillId         string `json:"bill_id"`
	Bill           Bill   `json:"bill"`
	Congress       int    `json:"congress"`
	Chamber        string `json:"chamber"`
	SourceType     string `json:"source_type"`
	LegislativeDay string `json:"legislative_day"`
	SchedudledAt   string `json:"scheduled_at"`
	Range          string `json:"range"`
	Context        string `json:"context"`
	URL            string `json:"url"`
}

type UpcomingBillResult struct {
	Result
	Results []UpcomingBill `json:"results"`
}

func UpcomingBills(criteria map[string]string) (*UpcomingBillResult, error) {
	l := UpcomingBillResult{}
	err := internal.GetURL(&l, congressRoot, criteria, "upcoming_bills")
	if err != nil {
		return nil, err
	}
	return &l, nil
}
