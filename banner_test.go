package main

import (
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

func TestGetCharLine(t *testing.T) {
	// first load the banner file.
	bannerLines, err := loadBanner("banner/standard.txt")
	if err != nil {
		t.Fatalf("could not load banner: %v", err)
	}

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
		result := getCharLines(bannerLines, tt.char)
		if len(result) != tt.expectedLines {
			t.Errorf("%s: expected %d lines but got %d", tt.name, tt.expectedLines, len(result))
		}
	}

}

func TestRenderWord(t *testing.T) {
	bannerLines, err := loadBanner("banner/standard.txt")
	if err != nil {
		t.Fatalf("couldn't load load banner banner: %v", err)
	}

	tests := []struct {
		name           string
		word           string
		expectedoutput string
	}{
		{"lwercase land", "land", "land"},
		{"uppercase ACER", "ACER", "ACER"},
		{"number 5", "23", "23"},
		{"space character", "land lord", "land lord"},
		{"last character tilde", "~", "~"},
	}

	for _, tt := range tests {
		result := renderWord(tt.word, bannerLines)
		if result != tt.expectedoutput {
			t.Errorf("%s: expected %s but got %s", tt.name, tt.expectedoutput, result)
		}
	}
}
