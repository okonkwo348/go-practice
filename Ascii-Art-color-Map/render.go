package main

import "strings"

func RenderLine(s string, substr string, banner map[rune][]string) []string {
	var build strings.Builder
	var result []string

	position := Position(s, substr)

	for i := 0; i < 8; i++ {
		for posCh, ch := range s {
			if isColor(posCh, position, len(substr)) {
				lines := banner[ch]
				build.WriteString("\033[33m" + lines[i] + "\033[33m")
			} else {
				lines := banner[ch]
				build.WriteString(lines[i])
			}
		}
		result = append(result, build.String())
		build.Reset()

	}
	return result
}
