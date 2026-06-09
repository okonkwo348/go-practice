package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Punc(s string) string {
	var b strings.Builder

	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " && unicode.IsPunct(rune(s[i+1])) {
			b.WriteString(string(s[i+1]))
			i++
			continue

		}

		b.WriteString(string(s[i]))
	}
	return b.String()
}

func main() {
	fmt.Println(Punc("Hello , world !"))
}
