package main

import "strings"

func renderWords(text string, bannerLines []string) []string {
	words := strings.Split(text, " ")
	var result []string

	for _, word := range words {
		art := RenderWord(word, bannerLines)
		result = append(result, art)

	}

	return result
}
