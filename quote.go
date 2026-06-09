package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Quote(s string) string {
	p := strings.Fields(s)
	k := strings.Join(p, " ")
	var b strings.Builder
	for i := 0; i < len(k); i++ {
		if string(k[i]) == " " && unicode.IsPunct(rune(k[i-1])) {
			b.WriteString(string(k[i+1]))
			i++
			continue
		}

		if string(k[i]) == " " && unicode.IsPunct(rune(k[i+1])) {
			b.WriteString(string(k[i+1]))
			i++
			continue
		}

		b.WriteString(string(k[i]))
	}

	return b.String()
}

func main() {
	fmt.Println(Quote(" '                   awasome      '"))
	fmt.Println(Quote(" '   hello world     '"))
}
