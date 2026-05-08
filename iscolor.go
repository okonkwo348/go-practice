package main

func isColor(i int, position []int, subLen int) bool {
	for _, pos := range position {
		if i >= pos && i < pos+subLen {
			return true
		}
	}
	return false
}
