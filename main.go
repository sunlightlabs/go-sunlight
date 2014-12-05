package main

import (
	"./openstates"
	"fmt"
)

func main() {
	leg := openstates.GetLegislator("COL000036")
	fmt.Printf(leg.FirstName + "\n")
	fmt.Printf(leg.LastName + "\n")
	for i := range leg.Sources {
		source := leg.Sources[i]
		fmt.Printf(source.Url)
	}
}
