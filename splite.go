package main

import (
	"fmt"
	"strings"
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

func main() {
	fmt.Println(Split("Hello, world! How are you?"))
	fmt.Println(Split2("Hello, world! How are you?"))
}
