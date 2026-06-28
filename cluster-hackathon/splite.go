package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Split(s string) []string {
	return strings.Fields(s)
}

func Split2(s string) []string {
	c := ""

	for _, ch := range s {
		if strings.ContainsAny(string(ch), ",.?!") {
			c += " " + string(ch)
			continue
		}
		c += string(ch)
	}

	return strings.Fields(c)
}

func Split3(s string) []string {
	var b strings.Builder

	for _, ch := range s {
		if unicode.IsPunct(ch) {
			b.WriteString(" " + string(ch))
			continue
		}
		b.WriteString(string(ch))
	}
	return strings.Fields(b.String())
}

func main() {
	fmt.Println(Split("Hello, world! How are you?"))
	fmt.Println(Split3("Hello, world! How are you?"))
}
