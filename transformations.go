package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// handleHex convert hex string (base 16) to decimal string (base 10).
func handleHex(s string) string {
	val, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return s // If it's not valid hex (like "ZZ"), just return it as is
	}
	return strconv.FormatInt(val, 10)
}

// handleBin convert binary string (base) to decimal string (base 10)
func handleBin(s string) string {
	val, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return s // If it's not valid hex (like "ZZ"), just return it as is
	}
	return strconv.FormatInt(val, 10)
}

// Function that convert a word to Title case
func capitalize(a string) string {
	s := strings.ToLower(a)
	if len(s) == 0 {
		return " "
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// function to handle Punctuation
func FixPunctuation(text string) string {
	reBefore := regexp.MustCompile(`\s+([,.;:?!])`)
	text = reBefore.ReplaceAllString(text, "$1")

	reAfter := regexp.MustCompile(`([,.;:?!])([a-zA-Z0-9])`)
	text = reAfter.ReplaceAllString(text, "$1 $2")

	return strings.TrimSpace(text)
}

// function to handle Quotation in an article
func FixQuotes(text string) string {
	reOpen := regexp.MustCompile(`'\s+([a-zA-Z0-9])`)
	text = reOpen.ReplaceAllString(text, "'$1")

	reClose := regexp.MustCompile(`([a-zA-Z0-9])\s+'`)
	text = reClose.ReplaceAllString(text, "$1'")

	return strings.TrimSpace(text)

}

// Function to handle Vowels in an article
func FixArticles(s []string) string {
	// s := strings.Fields(a)
	for i := 0; i < len(s); i++ {

		// Get lowercase versions for easy checking
		word := strings.ToLower(s[i])

		// 2. We only care if the current word is "a" or "an"
		if word == "an" || word == "a" {
			next := s[i+1]

			// Check if the NEXT word starts with a vowel or 'h'
			// We use nextWord[0] to look at the first character
			if strings.ContainsAny(string(next[0]), "aeiouhAEIOUH") {
				if s[i] == "An" || s[i] == "A" {
					s[i] = "An"
				} else {
					s[i] = "an"
				}

			} else {
				if s[i] == "a" || s[i] == "an" {
					s[i] = "a"
				} else {
					s[i] = "A"
				}
			}
		}
	}
	return strings.Join(s, " ")

}

func ApplyTransformations(input string) string {
	s := strings.Fields(input)
	var result []string // This will hold our final "clean" words

	for i := 0; i < len(s); i++ {
		word := s[i]

		// SENSOR: COMPLEX TAGS (up, n), (low, n), (cap, n)
		if i+1 < len(s) && strings.HasPrefix(word, "(") && strings.Contains(word, ",") {

			// THE EXTRACTION: "2)" -> "2"
			m := strings.TrimSuffix(s[i+1], ")")
			n, err := strconv.Atoi(m)

			if err == nil {

				// THE SAFETY GUARD: Don't look back further than we have words
				if n > len(result) {
					n = len(result)
				}

				for j := 1; j <= n; j++ {
					targetIndex := len(result) - j
					if strings.HasPrefix(word, "(up,") {
						result[targetIndex] = strings.ToUpper(result[targetIndex])
					} else if strings.HasPrefix(word, "(low,") {
						result[targetIndex] = strings.ToLower(result[targetIndex])
					} else if strings.HasPrefix(word, "(cap,") {
						result[targetIndex] = capitalize(result[targetIndex])
					}
				}
				i++
				continue
			}
		}

		// Use the switch for all single-word tags
		switch word {
		case "(hex)":
			if len(result) > 0 {
				last := len(result) - 1
				result[last] = handleHex(result[last])
			}
			// No append here, so (hex) is "deleted"

		case "(bin)":
			if len(result) > 0 {
				last := len(result) - 1
				result[last] = handleBin(result[last])
			}
			// No append here, so (bin) is "deleted"

		case "(up)":
			if len(result) > 0 {
				last := len(result) - 1
				result[last] = strings.ToUpper(result[last])
			}

		case "(low)":
			if len(result) > 0 {
				last := len(result) - 1
				result[last] = strings.ToLower(result[last])
			}
			// No append here, so (low) is "deleted"

		case "(cap)":
			if len(result) > 0 {
				last := len(result) - 1
				result[last] = capitalize(result[last])
			}
			// No append here, so (cap) is "deleted"

		default:
			result = append(result, word)

		}

	}

	//  Fix "A" vs "An" while still in slice form
	FixArticles(result)

	// Join into a single string for final spacing fixes
	finalText := strings.Join(result, " ")

	// Fix Quotes FIRST (so punctuation can still find the words)
	finalText = FixQuotes(finalText)

	// Fix Punctuation LAST
	finalText = FixPunctuation(finalText)

	return finalText
}
