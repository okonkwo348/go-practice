package main

import (
	"testing"
)

// Test for HandleHex
func TestHandleHex(t *testing.T) {
	result := handleHex("1E")
	expected := "30"

	if result != expected {
		t.Errorf("Expected %v, result %v", expected, result)
	}
}

// Test for handleBin
func TestHandleBin(t *testing.T) {
	result := handleBin("10")
	expected := "2"

	if result != expected {
		t.Errorf("Expected %v, result %v", expected, result)
	}
}

// Test for Capitalize
func TestCapitalize(t *testing.T) {
	result := capitalize("brEAd")
	expected := "Bread"

	if result != expected {
		t.Errorf("Expected %v, result %v", expected, result)
	}
}

// Test for punctuation
func TestFixpunctuation(t *testing.T) {
	result := FixPunctuation("I was thinking ... You were right. I was sitting over there ,and then BAMM !!")
	expected := "I was thinking... You were right. I was sitting over there, and then BAMM!!"

	if result != expected {
		t.Errorf("Expected %v, result %v", expected, result)
	}
}

// Test for Articlezg
func TestFixQuotes(t *testing.T) {
	result := FixQuotes("I am exactly how they describe me: ' awesome '")
	expected := "I am exactly how they describe me: 'awesome'"

	if result != expected {
		t.Errorf("Expected %v, result %v", expected, result)
	}
}

// Test FixArticles
func TestFixArticles(t *testing.T) {
	result := FixArticles([]string{"There it was. A amazing rock! An book"})
	expected := "There it was. A amazing rock! An book"

	if result != expected {
		t.Errorf("Expected %v, result %v", expected, result)
	}
}

// Test ApplyTransformations
func TestApplyTransformations(t *testing.T) {
	result := ApplyTransformations("it was THE (low, 1) best of times, it was the worst of times (up) IT WAS THE (low, 3) age of wisdom, it was the age of foolishness (cap, 6) it was the epoch of belief, it was the epoch of incredulity, it was the season of Light it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair. 1E (hex) files and 10 (bin) bytes.")
	expected := "it was the best of times, it was the worst of TIMES it was the age of wisdom, It Was The Age Of Foolishness it was the epoch of belief, it was the epoch of incredulity, it was the season of Light it was the season of darkness, it was the spring of hope, it was the winter of despair. 30 files and 2 bytes."

	if result != expected {
		t.Errorf("Expected %v, Result %v", expected, result)
	}

}
