package main

import (
	"testing"

	"github.com/jeroenrinzema/universal-name-translator/stapi"
)

func TestSpeciesLookup(t *testing.T) {
	uid := stapi.UID("000000000000")
	name := "Ambassador Spock"

	stapi.NewCharacterSearchMockServer(stapi.CharacterSearchResponse{
		Characters: []stapi.Character{
			{
				UID:  uid,
				Name: name,
			},
		},
	})

	stapi.NewBioMockServer(stapi.CharacterBioResponse{
		Character: stapi.Character{
			UID:  uid,
			Name: name,
			CharacterSpecies: stapi.CharacterSpecies{
				{
					UID:  uid,
					Name: "Human",
				},
				{
					UID:  uid,
					Name: "Vulkan",
				},
			},
		},
	})

	result, err := LookupCharacterSpecies("Ambassador Spock")
	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 2 {
		t.Errorf("unexpected ammount of character species returned. Expected 2 received %d\n", len(result))
	}
}
