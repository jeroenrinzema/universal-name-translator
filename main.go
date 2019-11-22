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

	LookupCharacterSpecies(input)
}
