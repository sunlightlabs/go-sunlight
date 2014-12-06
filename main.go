package main

import (
	"./openstates"
	"fmt"
)

func main() {
	legs := openstates.GetLegislators(map[string]string{
		"first_name": "john",
	})
	for _, leg := range *legs {
		fmt.Printf(leg.FirstName + " " + leg.LastName + "\n")
	}
}
