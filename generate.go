package main

import (
	"fmt"
	"strings"
)

func GeneratArt(word string, banner map[rune][]string) string {
	word = strings.ReplaceAll(word, "\n", "\\n")
	slice := Split(word)

	var result strings.Builder

	for i, word := range slice {
		if word == "" {
			if i == len(slice)-1 {
				continue
			}
			fmt.Println()
			continue
		}
		lines := RenderLine(word, banner)
		for _, line := range lines {
			result.WriteString(line + "\n")
		}
	}
	return result.String()
}
