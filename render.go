package main

import "strings"

func RenderWord(word string, bannerLines []string) string {
	var result strings.Builder

	for i := 0; i < 8; i++ {
		for _, ch := range word {
			start := int(ch-32)*9 + 1
			lines := bannerLines[start : start+8]
			result.WriteString(lines[i])
		}
		result.WriteString("\n")
	}

	return result.String()
}

func RenderAll(rows []string, bannerLines []string) string {
	var result strings.Builder
	for i, word := range rows {
		if word == "" {
			if i == len(rows)-1 {
				continue
			}
			result.WriteString("\n")
			continue
		}

		result.WriteString(RenderWord(word, bannerLines))
	}

	return result.String()
}
