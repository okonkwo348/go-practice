package main

import "strings"

func RenderLine(text string, banner map[rune][]string) []string {
	var result []string
	var contain strings.Builder
	for i := 0; i < 8; i++ {
		for _, char := range text {
			lines, ok := banner[char]
			if !ok {
				continue
			}
			contain.WriteString(lines[i])
		}
		result = append(result, contain.String())
		contain.Reset()
	}
	return result
}
