package main

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	// ch, err := ValidateInput(input)
	// if ch != 0 && err != nil {
	// 	return "unsupported char"
	// }

	var result strings.Builder
	input = strings.ReplaceAll(input, "\n", "\\n")
	splitedInput := SplitInput(input)

	for i, word := range splitedInput {
		if word == "" {
			if i == len(splitedInput)-1 {
				continue
			}
			result.WriteString("\n")
			continue
		}
		renderWord := RenderLine(word, banner)
		for i := 0; i < 8; i++ {

			result.WriteString(renderWord[i])
			result.WriteString("\n")
		}

	}
	return result.String()
}
