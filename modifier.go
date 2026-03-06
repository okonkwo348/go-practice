package main

import (
	"strings"
)

func applyUpper(input string) string {
	words := strings.Fields(input)
	var result []string // This will hold our final "clean" words

	for _, word := range words {
		if word == "(up)" {
			// 1. Check if there is a word before this one in our 'result' slice
			if len(result) > 0 {
				// 2. Uppercase the LAST item we added to 'result'
				lastIndex := len(result) - 1
				result[lastIndex] = strings.ToUpper(result[lastIndex])
			}
			// 3. NOTICE: We do NOT append "(up)" to the result slice.
			// This effectively "deletes" it!
		} else {
			// It's a normal word, just add it to the pile
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}
