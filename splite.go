package main

import (
	"fmt"
	"strings"
)

func Split(s string) []string {
	return strings.Fields(s)
}

func Split2(s string) []string {
	return strings.Split(s, "!")
}

func main() {
	fmt.Println(Split("Hello, world! How are you?"))
	fmt.Println(Split2("Hello, world! How are you?"))
}
