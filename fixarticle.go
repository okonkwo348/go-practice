package main

import (
	"fmt"
	"strings"
)

func fixArticles(words []string) []string {
	for i := 0; i < len(words); i++ {

		if words[i] == "an" || words[i] == "a" || words[i] == "An" || words[i] == "A" {
			next := words[i+1]
			if strings.ContainsAny(string(next[0]), "haeiouHAEIOU") {
				if words[i] == "An" || words[i] == "an" {
					if words[i] == "An" {
						words[i] = "A"
					} else {
						words[i] = "a"
					}
				} else if words[i] == "A" || words[i] == "a" {
					if words[i] == "A" {
						words[i] = "An"
					} else {
						words[i] = "an"
					}

				}
			}
		}
	}

	return words
}

func main() {
	fmt.Println(fixArticles([]string{"a", "egg", "an", "kar", "An", "louse"}))
}
