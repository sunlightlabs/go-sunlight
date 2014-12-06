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
	fmt.Printf(meta.Name + "\n")
}
