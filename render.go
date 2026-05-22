package main

import "strings"

func RenderLine(s string, banner map[rune][]string) []string {
	var contain strings.Builder

	var result []string
	for i := 0; i < 8; i++ {
		for _, ch := range s {
			lines := banner[ch]
			contain.WriteString(lines[i])
		}
		result = append(result, contain.String())
		contain.Reset()
	}
	return result
}
