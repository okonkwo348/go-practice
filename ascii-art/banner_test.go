package main

import (
	"strings"
	"testing"
)

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		expectError bool
	}{
		{"valid normal string", []string{"hello"}, false},
		{"valid string with space", []string{"hello world"}, false},
		{"empty string", []string{""}, false},
		{"invalid character", []string{"\x01"}, true},
		{"multiple valid rows", []string{"hello", "world"}, false},
		{"string with numbers", []string{"hello123"}, false},
	}

	for _, tt := range tests {
		_, err := validateInput(tt.input)
		if tt.expectError && err == nil {
			t.Errorf("%s: expected error but got none", tt.name)
		}

		if !tt.expectError && err != nil {
			t.Errorf("%s: unexpected error: %v", tt.name, err)
		}
	}
}

var testBannerLines []string

func TestMain(m *testing.M) {
	var err error
	testBannerLines, err = loadBanner("banner/standard.txt")
	if err != nil {
		panic("could not load banner")
	}
	m.Run() // runs all test
}

func TestGetCharLine(t *testing.T) {

	tests := []struct {
		name          string
		char          rune
		expectedLines int
	}{
		{"lowercase a", 'a', 8},
		{"uppercase A", 'A', 8},
		{"space character", ' ', 8},
		{"number 5", '5', 8},
		{"last character tilde", '~', 8},
	}

	for _, tt := range tests {
		result := getCharLines(testBannerLines, tt.char)
		if len(result) != tt.expectedLines {
			t.Errorf("%s: expected %d lines but got %d", tt.name, tt.expectedLines, len(result))
		}
	}

}

func TestRenderWord(t *testing.T) {

	tests := []struct {
		name          string
		word          string
		expectedLines int
	}{
		{"single char", "A", 8},
		{"word", "hello", 8},
		{"empty string", "", 8},
		{"with space", "hi there", 8},
	}

	for _, tt := range tests {
		result := renderWord(tt.word, testBannerLines)
		lines := strings.Split(result, "\n")
		// remove last empty element from trailing \n
		if lines[len(lines)-1] == "" {
			lines = lines[:len(lines)-1]
		}
		if len(lines) != tt.expectedLines {
			t.Errorf("%s: expected %d lines but got %d", tt.name, tt.expectedLines, len(lines))
		}
	}
}
