package main

func isColored(i int, positions []int, subLen int) bool {
	for _, pos := range positions {
		if i >= pos && i < pos+subLen {
			return true
		}
	}
	return false
}
