package main

import (
	"fmt"
	"strings"
)

func Split(s string) []string {
	return strings.Fields(s)
}

// func Split(s string) []string {
// 	return strings.Split(s, "!")
// }

func main() {
	fmt.Println(Split("Hello, world! How are you?"))
}
