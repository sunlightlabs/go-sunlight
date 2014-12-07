package main

import (
	"./openstates"
	"fmt"
)

func main() {
	legs, err := openstates.GetLegislatorsByLatLon(38.889730, -77.017626)
	if err != nil {
		panic(err.Error())
	}
	for _, leg := range *legs {
		fmt.Printf(leg.FullName + "\n")
	}
}
