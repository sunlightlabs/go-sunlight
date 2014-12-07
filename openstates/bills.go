package openstates

import (
	"../internal"
)

type Bill struct {
	BillId   string   `json:"bill_id"`
	Id       string   `json:"id"`
	Chamber  string   `json:"chamber"`
	Session  string   `json:"session"`
	State    string   `json:"state"`
	Subjects []string `json:"subjects"`
	Title    string   `json:"title"`
	Types    []string `json:"type"`
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
