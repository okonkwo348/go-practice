package main

import (
	"strings"
)

func GenerateArt(word string, subStr string, banner map[rune][]string) string {

	if word == "" {
		return ""
	}

	word = strings.ReplaceAll(word, "\n", "\\n")
	slice := Split(word)

	var result strings.Builder

	for i, word := range slice {
		if word == "" {
			if i == len(slice)-1 {
				continue
			}
			result.WriteString("\n")
			continue
		}
		lines := RenderLine(word, subStr, banner)
		for _, line := range lines {
			result.WriteString(line + "\n")
		}
	}
	return result.String()
}
