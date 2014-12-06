package main

import (
	"./openstates"
	"fmt"
)

func main() {
	legs := openstates.GetLegislators(map[string]string{
		"state": "ma",
	})
	for _, leg := range *legs {
		fmt.Printf(leg.FirstName + "\n")
	}
}
