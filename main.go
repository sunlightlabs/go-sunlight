package main

import (
	"./openstates"
	"fmt"
)

func main() {
	meta, err := openstates.GetMetadata("ma")
	if err != nil {
		panic(err.Error())
	}
	for _, term := range meta.Terms {
		fmt.Printf(term.Name + "\n")
	}
}
