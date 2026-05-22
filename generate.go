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
		if i < len(slice)-1 {
			result.WriteString("\n")
			continue
		}

		renderLines := RenderLine(word, banner)
		for _, lines := range renderLines {
			result.WriteString(lines + "\n")
		}
	}
	return result.String()
}
