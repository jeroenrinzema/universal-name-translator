package main

import "testing"

import "strings"

func TestUTF8UnicodeLookup(t *testing.T) {
	value := "a"
	length := 10

	characters := Characters{}
	expected := strings.Repeat(value, length)

	for i := 0; i <= length; i++ {
		value := strings.Repeat(value, i)
		characters[value] = value
	}

	result, _ := characters.UTF8UnicodeLookup(expected)
	if len(result) > 1 {
		t.Error("Multiple results are returned when one is expected")
	}

	if result[0] != expected {
		t.Errorf("Unexpected result returned. Expected: %s, received %s\n", expected, result[0])
	}
}

func TestUTF8LookupFirstMatch(t *testing.T) {
	value := "a"
	total := 10

	characters := Characters{}
	input := strings.Repeat(value, total/2)

	for i := 0; i <= total; i++ {
		value := strings.Repeat(value, i)
		characters[value] = value
	}

	match, _ := UTF8LookupFirstMatch(input, characters, 0)
	if match != input {
		t.Errorf("Unexpected result returned. Expected: %s, received %s\n", input, match)
	}
}
