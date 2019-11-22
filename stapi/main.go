package stapi

import "errors"

// ErrNotFound is returned when a search querie result contains no results
var ErrNotFound = errors.New("not found")

// UID represents a unique identifier of an stored entity
type UID string

// CharacterSpecies holds possibly multiple character species for a Star Trek character
type CharacterSpecies []Species

// String returns the character species name(s) seperated by a comma
func (species CharacterSpecies) String() (result string) {
	for index, species := range species {
		if index != 0 {
			result = result + " ,"
		}

		result = result + species.Name
	}

	return result
}

// Species represents a Star Trek character species
type Species struct {
	UID  UID    `json:"uid"`
	Name string `json:"name"`
}

// Character represents a Star Trek character
type Character struct {
	UID              UID              `json:"uid"`
	Name             string           `json:"name"`
	CharacterSpecies CharacterSpecies `json:"characterSpecies"`
}
