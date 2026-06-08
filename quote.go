package main

import (
	"fmt"
	"regexp"
)

func Quote(s string) string {
	last := regexp.MustCompile("'*s/([^'])*s/'")
	s = last.ReplaceAllString(s, "'$1'")

	return s
}

func main() {
	fmt.Println(Quote(" ' awasome '"))
}
