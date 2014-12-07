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
	BillId     string   `json:"bill_id"`
	Id         string   `json:"id"`
	AllIds     []string `json:"all_ids"`
	Companions []string `json:"companions"`

	AlternateTitles []string `json:"alternate_titles"`
	Title           string   `json:"title"`

	Country string `json:"country"`
	Level   string `json:"level"`
	Session string `json:"session"`
	State   string `json:"state"`

	Actions  []BillAction  `json:"actions"`
	Sponsors []BillSponsor `json:"sponsors"`
	Chamber  string        `json:"chamber"`

	Subjects []string `json:"subjects"`
	Types    []string `json:"type"`

	Sources []Source `json:"sources"`
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
