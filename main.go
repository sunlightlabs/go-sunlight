package main

import (
	"./openstates"
	"fmt"
)

func main() {
	// bill, err := openstates.GetBill("wy", "2013", "HB 149")
	bills, err := openstates.GetBills(map[string]string{
		"q": "beer",
	})
	if err != nil {
		panic(err.Error())
	}
	for _, bill := range *bills {
		fmt.Printf(bill.Title + "\n")
	}
}
