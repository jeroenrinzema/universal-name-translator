package stapi

import (
	"testing"
)

func TestCharacterSearch(t *testing.T) {
	uid := UID("000000000000")
	name := "Ambassador Spock"

	NewCharacterSearchMockServer(CharacterSearchResponse{
		Characters: []Character{
			{
				UID:  uid,
				Name: name,
			},
		},
	})

	result, err := CharacterSearch("Ambassador Spock")
	if err != nil {
		t.Fatal(err)
	}

	if result.UID != uid {
		t.Errorf("unexpected character name. Expected %s, received %s\n", uid, result.UID)
	}

	if result.Name != name {
		t.Errorf("unexpected character name. Expected %s, received %s\n", name, result.Name)
	}
}

func TestCharacterBIOLookup(t *testing.T) {
	uid := UID("000000000000")
	name := "Ambassador Spock"

	NewBioMockServer(CharacterBioResponse{
		Character: Character{
			UID:  uid,
			Name: name,
			CharacterSpecies: CharacterSpecies{
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

	result, err := CharacterBio(uid)
	if err != nil {
		t.Fatal(err)
	}

	if result.UID != uid {
		t.Errorf("unexpected character name. Expected %s, received %s\n", uid, result.UID)
	}

	if result.Name != name {
		t.Errorf("unexpected character name. Expected %s, received %s\n", name, result.Name)
	}

	if len(result.CharacterSpecies) != 2 {
		t.Errorf("unexpected character species returned. Expected 2, received %d\n", len(result.CharacterSpecies))
	}
}
