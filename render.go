package main

import "strings"

func RenderLine(text string, banner map[rune][]string) []string {
	var result []string
	var contain strings.Builder
	for i := 0; i < 8; i++ {
		for _, char := range text {
			contain.WriteString(banner[char][i])
		}
		result = append(result, contain.String())
		contain.Reset()
	}
	return result
}
