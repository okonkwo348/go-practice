package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Upper(s string) string {
	return strings.ToUpper(s)
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func onWordCap(a string) string {
	s := strings.ToLower(a)
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func Cap(s string) string {
	a := strings.Fields(s)
	for i := 0; i < len(a); i++ {
		a[i] = onWordCap(a[i])
	}
	return strings.Join(a, " ")
}

func Title(a string) string {
	s := strings.Fields(a)
	for i := 0; i < len(s); i++ {
		if len(s[i]) > 3 || i == 0 {
			s[i] = onWordCap(s[i])
		}
	}
	return strings.Join(s, " ")
}
func main() {
	fmt.Println(Title("my fAtheR is go to cHurch"))
}
