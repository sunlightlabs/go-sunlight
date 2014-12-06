package main

import (
	"./openstates"
	"fmt"
)

func main() {
	legs, err := openstates.GetLegislators(map[string]string{
		"first_name": "john",
	})
	if err != nil {
		panic(err.Error())
	}
	for _, leg := range *legs {
		fmt.Printf(leg.FirstName + " " + leg.LastName + "\n")
	}
}
