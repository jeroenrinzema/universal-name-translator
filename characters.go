package main

import (
	"errors"
	"strings"
)

// ErrUntranslatableCharacters is returned when a untranslatable character is encountered
var ErrUntranslatableCharacters = errors.New("untranslatable characters encountered")

// Characters represents a lookup table of UTF-8 characters and their foreign unicode representative
type Characters map[string]string

// UTF8UnicodeLookup looks up the unicode character representations of the given UTF-8 string.
func (characters Characters) UTF8UnicodeLookup(input string) (result []string, err error) {
	for {
		match, value := UTF8LookupFirstMatch(input, characters, 0)

		if len(match) == 0 {
			return nil, ErrUntranslatableCharacters
		}

		result = append(result, value)
		input = input[len(match):]

		if len(input) == 0 {
			break
		}
	}

	return result, nil
}

// UTF8LookupFirstMatch looks up the first character representation that has the highes match rate
// for the given input. The matched key and representing value are returned.
func UTF8LookupFirstMatch(input string, characters Characters, position int) (match string, result string) {
	matches := Characters{}
	input = strings.ToLower(input)

	for key, value := range characters {
		// TODO: in some character sets capitalised letters could represent different characters.
		// For now we ignore this use case and transform the input and to lower case characters
		key = strings.ToLower(key)

		if len(key) < position+1 {
			continue
		}

		if len(input) < len(key) {
			continue
		}

		if string(input[position]) != string(key[position]) {
			continue
		}

		matches[key] = value

		match = key
		result = value
	}

	if len(matches) > 1 {
		return UTF8LookupFirstMatch(input, matches, position+1)
	}

	if match != input[:len(match)] {
		return "", ""
	}

	return match, result
}

// Klingon holds the klingon alphabet and its UTF-8 representing characters
var Klingon = Characters{
	"a":   "0xF8D0",
	"b":   "0xF8D1",
	"ch":  "0xF8D2",
	"D":   "0xF8D3",
	"e":   "0xF8D4",
	"gh":  "0xF8D5",
	"H":   "0xF8D6",
	"I":   "0xF8D7",
	"j":   "0xF8D8",
	"l":   "0xF8D9",
	"m":   "0xF8DA",
	"n":   "0xF8DB",
	"ng":  "0xF8DC",
	"o":   "0xF8DD",
	"p":   "0xF8DE",
	"q":   "0xF8DF",
	"Q":   "0xF8E0",
	"r":   "0xF8E1",
	"S":   "0xF8E2",
	"t":   "0xF8E3",
	"tlh": "0xF8E4",
	"u":   "0xF8E5",
	"v":   "0xF8E6",
	"w":   "0xF8E7",
	"y":   "0xF8E8",
	"'":   "0xF8E9",
	"1":   "0xF8F0",
	"2":   "0xF8F1",
	"3":   "0xF8F2",
	"4":   "0xF8F3",
	"5":   "0xF8F4",
	"6":   "0xF8F5",
	"7":   "0xF8F6",
	"8":   "0xF8F7",
	"9":   "0xF8F8",
	"0":   "0xF8F9",
	".":   "0xF8FD",
	",":   "0xF8FE",
	" ":   "0x0020",
}
