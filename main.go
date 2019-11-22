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

	translate(input)
	species(input)
}

func translate(input string) {
	result, err := Klingon.UTF8UnicodeLookup(input)
	if err != nil {
		fmt.Println("input contains untranslatable characters.")
		return
	}

	fmt.Println(strings.Join(result, " "))
}

func species(input string) {
	species, err := LookupCharacterSpecies(input)
	if err != nil {
		fmt.Println("unable to lookup species.")
		return
	}

	fmt.Println(species)
}
