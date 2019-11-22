package stapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// CharacterSearchURL holds the STAPI API endpoint for a character search
var CharacterSearchURL = "http://stapi.co/api/v1/rest/character/search"

// CharacterBioURL holds the STAPI API endpoint for a character bio lookup
var CharacterBioURL = "http://stapi.co/api/v1/rest/character"

// CharacterSearchResponse represents the character search JSON response body
type CharacterSearchResponse struct {
	Characters []Character `json:"characters"`
}

// CharacterBioResponse represents the character bio response body
type CharacterBioResponse struct {
	Character Character `json:"character"`
}

// CharacterSearch attempts to preform a search querie for the given character name.
// The first character returned inside the HTTP response is returned
func CharacterSearch(name string) (Character, error) {
	form := url.Values{}
	form.Set("name", name)

	bb := strings.NewReader(form.Encode())
	res, err := http.Post(CharacterSearchURL, "application/x-www-form-urlencoded", bb)
	if err != nil {
		return Character{}, err
	}

	defer res.Body.Close()

	response := CharacterSearchResponse{}
	json.NewDecoder(res.Body).Decode(&response)

	if len(response.Characters) == 0 {
		return Character{}, ErrNotFound
	}

	return response.Characters[0], nil
}

// CharacterBio attempts to fetch additional bio information for the given character uid
func CharacterBio(uid UID) (Character, error) {
	endpoint, err := url.Parse(CharacterBioURL)
	if err != nil {
		return Character{}, err
	}

	query := endpoint.Query()
	query.Add("uid", string(uid))

	endpoint.RawQuery = query.Encode()

	res, err := http.Get(endpoint.String())
	if err != nil {
		return Character{}, err
	}

	defer res.Body.Close()

	response := CharacterBioResponse{}
	json.NewDecoder(res.Body).Decode(&response)

	return response.Character, nil
}
