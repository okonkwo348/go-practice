package main

import "strings"

func Position(text string, substring string) []int {
	start := 0
	var result []int

	for {
		idx := strings.Index(text[start:], substring)

		if idx == -1 {
			break
		}

		foundAt := idx

		result = append(result, foundAt+len(substring))

		start = foundAt + 1
	}

	return result
}
