package main

import "strings"

func findAllPositions(text, substr string) []int {
	var positions []int
	start := 0
	for {
		idx := strings.Index(text[start:], substr)

		if idx == -1 {
			break
		}

		pos := start + idx
		positions = append(positions, pos)

		start = pos + 1

	}
	return positions
}
