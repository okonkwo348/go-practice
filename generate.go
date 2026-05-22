package main

import "strings"

func GenerateArt(s string, banner map[rune][]string) string {

	if s == "" {
		return ""
	}
	s = strings.ReplaceAll(s, "\n", "\\n")

	slice := SplitInput(s)

	var result strings.Builder

	for i, word := range slice {
		if word == "" {
			if i == len(slice)-1 {
				continue
			}
			result.WriteString("\n")
			continue
		}

		renderLines := RenderLine(word, banner)

		for i := 0; i < 8; i++ {
			result.WriteString(renderLines[i])
			result.WriteString("\n")
		}
	}
	return result.String()
}
