package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	args := os.Args[1:]
	input := strings.Join(args, " ")

	result := Klingon.UTF8UnicodeLookup(input)
	fmt.Println(strings.Join(result, " "))

	species, err := LookupCharacterSpecies(input)
	if err != nil {
		fmt.Println("unable to lookup species.")
		return
	}

	fmt.Println(species)
}
