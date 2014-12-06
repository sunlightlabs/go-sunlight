package main

import (
	"./openstates"
	"fmt"
)

func main() {
	meta, err := openstates.GetMetadataList()
	if err != nil {
		panic(err.Error())
	}
	for _, state := range *meta {
		fmt.Printf(state.Name + "\n")
	}
}
