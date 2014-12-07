package main

import (
	"./openstates"
	"fmt"
)

func main() {
	bill, err := openstates.GetBill("wy", "2013", "HB 149")
	if err != nil {
		panic(err.Error())
	}
	for _, sponsor := range bill.Sponsors {
		fmt.Printf(sponsor.Name + "\n")
	}
}
