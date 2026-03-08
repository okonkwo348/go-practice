package main

import (
	"strconv"
	"strings"
)

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

			} else {
				result = append(result, word)
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

			} else {
				result = append(result, word)
			}
		}

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
				result[preWord] = strings.ToTitle(result[preWord])
			}

		default:
			// It's a normal word, just add it to the pile
			result = append(result, word)

		}
	}
	return strings.Join(result, " ")
}
