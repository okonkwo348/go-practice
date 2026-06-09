package main

import (
	"fmt"
	"regexp"
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

func Punc1(s string) string {
	punc := regexp.MustCompile(`\s*({a-zA-Z}) ({,"?!})`)
	return punc.ReplaceAllString(s, "$1$2")
}

func main() {
	fmt.Println(Punc("Hello , world !"))
	fmt.Println(Punc1("Hello , world !"))
}
