package stapi

import "errors"

// ErrNotFound is returned when a search querie result contains no results
var ErrNotFound = errors.New("not found")

// UID represents a unique identifier of an stored entity
type UID string

// Species represents a Star Trek character species
type Species struct {
	UID  UID    `json:"uid"`
	Name string `json:"name"`
}

// Character represents a Star Trek character
type Character struct {
	UID              UID       `json:"uid"`
	Name             string    `json:"name"`
	CharacterSpecies []Species `json:"characterSpecies"`
}
