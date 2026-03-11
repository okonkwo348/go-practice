package main

import (
	"strconv"
	"strings"
	"unicode"
)

// manual capitalize function using standard library

func capitalize(s string) string {
	if len(s) == 0 {
		return " "
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func applyUpper(input string) string {
	words := strings.Fields(input)
	result := []string{} // This will hold our final "clean" words

	for i := 0; i < len(result); i++ {
		word := words[i]

		// modification into Hexadecimal
		if word == "(hex)" {
			if len(result) > 0 {
				preStr := len(result) - 1

				// conversion of string to base hexadecimal
				hexStr, err := strconv.ParseInt(result[preStr], 16, 64)

				// continue the iteration if error is found
				if err != nil {
					continue
				}

				// conversion of hexadeimal to string
				result[preStr] = strconv.FormatInt(hexStr, 10)
				// (hex) is deleted

			}
		}

		// modification into Binary
		if word == "(bin)" {
			if len(result) > 0 {
				preStr := len(result) - 1

				// conversion of string to binary
				hexStr, err := strconv.ParseInt(result[preStr], 2, 64)

				// continue the iteration if error is found
				if err != nil {
					continue
				}

				// convert of binary to string
				result[preStr] = strconv.FormatInt(hexStr, 10)

			}
		}

		//  SIMPLE TAGS (up), (low), (cap)
		// We use a switch because we are checking for EXACT matches here.
		switch word {

		// Modification into Uppercase
		case "(up)":
			//  Check if there is a word before this one in our 'result' slice
			if len(result) > 0 {
				// 2. Uppercase the LAST item we added to 'result'
				preWord := len(result) - 1
				result[preWord] = strings.ToUpper(result[preWord])
			}
			// NOTICE: We do NOT append "(up)" to the result slice.
			// This effectively "deletes" it

		// Modification into Lowercase
		case "(low)":
			if len(result) > 0 {
				preWord := len(result) - 1
				result[preWord] = strings.ToLower(result[preWord])
			}

		// Modification into Titlecase or capitalization
		case "cap":
			if len(result) > 0 {
				preWord := len(result) - 1
				result[preWord] = capitalize(result[preWord])
			}

		default:
			// ONLY append if it's NOT a tag
			result = append(result, word)

		}

		// SENSOR: COMPLEX TAGS (up, n), (low, n), (cap, n)
		if i+1 < len(words) && (strings.HasPrefix(word, "(up,") || strings.HasPrefix(word, "(low,") || strings.HasPrefix(word, "(cap,")) {

			// THE EXTRACTION: "2)" -> "2"
			numberStr := strings.TrimSuffix(words[i+1], ")")
			n, err := strconv.Atoi(numberStr)

			if err == nil {
				// THE SAFETY GUARD: Don't look back further than we have words

				if n > len(result) {
					n = len(result)
				}

				// THE REVERSE LOOP: Count back 'n' times
				for j := 1; j <= n; j++ {
					targetIndex := len(result) - j

					// Determine WHICH transformation to apply
					if strings.HasPrefix(word, "(up,") {
						result[targetIndex] = strings.ToUpper(result[targetIndex])
					} else if strings.HasPrefix(word, "(low,") {
						result[targetIndex] = strings.ToLower(result[targetIndex])
					} else if strings.HasPrefix(word, "(cap,") {
						result[targetIndex] = capitalize(result[targetIndex])
					}
				}
				i++      // Skip the number word
				continue // Skip appending
			}
		}
		result = append(result, word)
	}
	return strings.Join(result, " ")

}
