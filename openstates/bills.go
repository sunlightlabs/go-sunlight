package openstates

import (
	"../internal"
)

type BillAction struct {
	Action string   `json:"action"`
	Actor  string   `json:"actor"`
	Types  []string `json:"type"`
}

type BillSponsor struct {
	LegId string `json:"leg_id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

type BillVersion struct {
	DocId    string `json:"doc_id"`
	Mimetype string `json:"mimetype"`
	Name     string `json:"name"`
	Url      string `json:"url"`
}

type Bill struct {
	BillId          string       `json:"bill_id"`
	Id              string       `json:"id"`
	AllIds          []string     `json:"all_ids"`
	AlternateTitles []string     `json:"alternate_titles"`
	Companions      []string     `json:"companions"`
	Country         string       `json:"country"`
	Sources         []Source     `json:"sources"`
	Level           string       `json:"level"`
	Session         string       `json:"session"`
	Actions         []BillAction `json:"actions"`
	Sponsors        []Sponsor    `json:"sponsors"`
	Chamber         string       `json:"chamber"`
	Session         string       `json:"session"`
	State           string       `json:"state"`
	Subjects        []string     `json:"subjects"`
	Title           string       `json:"title"`
	Types           []string     `json:"type"`
}

func GetBill(state, session, billId string) (*Bill, error) {
	t := &Bill{}
	err := internal.GetURL(t, openstatesRoot, map[string]string{}, "bills", state, session, billId)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func GetBills(criteria map[string]string) (*[]Bill, error) {
	t := &[]Bill{}
	err := internal.GetURL(t, openstatesRoot, criteria, "bills")
	if err != nil {
		return nil, err
	}
	return t, nil
}

func GetBillById(bigId string) (*Bill, error) {
	t := &Bill{}
	err := internal.GetURL(t, openstatesRoot, map[string]string{}, "bills", bigId)
	if err != nil {
		return nil, err
	}
	return t, nil
}
