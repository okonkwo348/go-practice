package main

import "strings"

// For each character, store its 8 lines
func RenderWord(s string, bannerLines []string) string {
	var result strings.Builder

	allCharLines := [][]string{}
	for _, char := range s {
		linesOfEachChar := renderChar(bannerLines, char)
		allCharLines = append(allCharLines, linesOfEachChar)
	}

	for i := 0; i < 8; i++ {
		for _, charLine := range allCharLines {
			result.WriteString(charLine[i])
		}
		result.WriteString("\n")
	}
	return result.String()
}
