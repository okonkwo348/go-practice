package main

import (
	"fmt"
	"strings"
)

func fixArticles(words []string) []string {
	for i := 0; i < len(words); i++ {

		if words[i] == "A" || words[i] == "a" {
			next := words[i+1]
			if i+1 < len(words) && strings.ContainsAny(string(next[:1]), "haeiouHAEIOU") {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}

			}
		}

		if words[i] == "An" || words[i] == "an" {
			next := words[i+1]
			if i+1 < len(words) && !strings.ContainsAny(string(next[0]), "haeiouHAEIOU") {

				if words[i] == "An" {
					words[i] = "A"
				} else {
					words[i] = "a"
				}

			}
		}
	}

	return words

}

func main() {
	fmt.Println(fixArticles([]string{"a", "egg", "an", "kar", "An", "louse", "A", "unto", "An", "goat"}))
}
