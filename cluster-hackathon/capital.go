package main

import (
	"fmt"
	"strings"
	"unicode"
)

func capitalize(s string) string {
	s = strings.ToLower(s)
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)

}

func capitalizeN(words []string, n int) []string {

	for i := 0; i < len(words); i++ {
		if i >= n-1 {
			words[i] = strings.Title(words[i])

		}
	}

	return words
}

func main() {
	fmt.Println(capitalize("emma"))
	fmt.Println(capitalizeN([]string{"this", "is", "my", "world", "so", "get", "that"}, 4))
}
