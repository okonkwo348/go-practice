package main

func StackTwo(top []string, bottom []string) []string {

	n := len(top) + len(bottom)
	result := make([]string, n)

	copy(result[:len(top)], top)

	copy(result[len(top):], bottom)

	return result
}

func StackAll(blocks [][]string) []string {
	result := []string{}

	for _, block := range blocks {
		StackTwo(result, block)
	}

	return result
}
