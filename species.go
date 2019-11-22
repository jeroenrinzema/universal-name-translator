package main

import (
	"errors"

	"github.com/jeroenrinzema/universal-translator/stapi"
)

// ErrNoSpecies is returned when no species are found for the given character
var ErrNoSpecies = errors.New("given character is not part of any species")

// LookupCharacterSpecies attempts to look up the species as a string of the given character.
// If a character is part of multiple species all species are joined and seperated by a comma
func LookupCharacterSpecies(name string) (result stapi.CharacterSpecies, err error) {
	character, err := stapi.CharacterSearch(name)
	if err != nil {
		return nil, err
	}

	bio, err := stapi.CharacterBio(character.UID)
	if err != nil {
		return nil, err
	}

	if len(bio.CharacterSpecies) == 0 {
		return nil, ErrNoSpecies
	}

	return bio.CharacterSpecies, nil
}
