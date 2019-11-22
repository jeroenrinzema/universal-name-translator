package stapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// NewCharacterSearchMockServer constructs a new HTTP server that responds to bio requests
func NewCharacterSearchMockServer(result CharacterSearchResponse) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("ContentType", "application/json")
		w.WriteHeader(http.StatusOK)

		bb, _ := json.Marshal(result)
		w.Write(bb)
	}))

	CharacterSearchURL = server.URL

	return server
}

// NewBioMockServer constructs a new HTTP server that responds to bio requests
func NewBioMockServer(character CharacterBioResponse) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("ContentType", "application/json")
		w.WriteHeader(http.StatusOK)

		bb, _ := json.Marshal(character)
		w.Write(bb)
	}))

	CharacterBioURL = server.URL

	return server
}
