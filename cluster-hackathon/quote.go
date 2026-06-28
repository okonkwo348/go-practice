package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Quote(s string) string {
	p := strings.Fields(s)
	k := strings.Join(p, " ")
	var b strings.Builder
	for i := 0; i < len(k); i++ {
		if string(k[i]) == " " && strings.ContainsAny(string((k[i-1])), "'") {
			b.WriteString(string(k[i+1]))
			i++
			continue
		}

		if string(k[i]) == " " && strings.ContainsAny(string((k[i+1])), "'") {
			b.WriteString(string(k[i+1]))
			i++
			continue
		}

		b.WriteString(string(k[i]))
	}

	return b.String()
}

func Quote2(s string) string {
	last := regexp.MustCompile(`'\s*([^']*?)\s*'`)
	return last.ReplaceAllString(s, "'$1'")
}

func main() {
	fmt.Println(Quote2(" '                   awasome      '"))
	fmt.Println(Quote2("'   hello , world   '  "))
}
