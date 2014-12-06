package main

import (
	"./openstates"
	"fmt"
)

func main() {
	legs, err := openstates.GetLegislators(map[string]string{
		"first_name": "john",
		"fields":     "roles,first_name,last_name",
	})
	if err != nil {
		panic(err.Error())
	}
	for _, leg := range *legs {
		fmt.Printf(leg.FirstName + " " + leg.LastName + "\n")
		for _, role := range leg.Roles {
			if role.District != "" {
				fmt.Printf("  " + role.Term + ", " + role.District + " (" + role.Chamber + ")\n")
			}
		}
	}
}
