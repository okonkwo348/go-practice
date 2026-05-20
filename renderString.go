package main

import "strings"

func renderString(text string, bannerLines []string) string {
	words := strings.Split(text, "\\n")
	var result strings.Builder

	for _, word := range words {
		art := RenderWord(word, bannerLines)
		result.WriteString(art)

	}

	return result.String()
}
