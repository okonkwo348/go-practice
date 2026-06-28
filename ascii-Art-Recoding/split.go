package main

import "strings"

func SplitInput(input string) []string {
	slice := strings.Split(input, "\\n")
	return slice

}
