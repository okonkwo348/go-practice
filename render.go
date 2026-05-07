package main

import "strings"

func RenderLine(s string, banner map[rune][]string) []string {
	var build strings.Builder
	var result []string
	for i := 0; i < 8; i++ {
		for _, ch := range s {
			lines := banner[ch]
			build.WriteString(lines[i])
		}
		result = append(result, build.String())
		build.Reset()

	}
	return result
}
