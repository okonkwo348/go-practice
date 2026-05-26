package main

import "strings"

func renderWord(word string, bannerLines []string) string {
	var result strings.Builder
	for i := 0; i < 8; i++ {
		for _, char := range word {
			line := getCharLines(bannerLines, char)
			result.WriteString(line[i])
		}
		result.WriteString("\n")
	}

	return result.String()
}
