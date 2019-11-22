package main

import (
	"github.com/jeroenrinzema/universal-translator/stapi"
)

// LookupCharacterSpecies attempts to look up the species as a string of the given character.
func LookupCharacterSpecies(name string) (string, error) {
	_, err := stapi.CharacterSearch(name)
	if err != nil {
		return "", err
	}

	return "", nil
}
