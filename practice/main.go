package main

import (
	"fmt"
	"unicode"
)

func cameltoSnake(s string) string {
	n := []rune(s)
	result := []rune{}
	for i := 0; i < len(n); i++ {
		if n[i] == unicode.ToUpper(n[i]) {
			n[i] = unicode.ToLower(n[i])
			result = append(result, '_')
			result = append(result, n[i])
			continue
		}
		result = append(result, n[i])
	}
	return string(result)
}

// hggh
func snaketoCamel(s string) string {
	n := []rune(s)
	result := []rune{}
	for i := 0; i < len(n); i++ {
		if n[i] == '_' {
			n[i+1] = unicode.ToUpper(n[i+1])
			continue
		}
		result = append(result, n[i])
	}
	return string(result)
}
func main() {
	fmt.Println(cameltoSnake("nameEmmaOkonkwo"))
	//Expected Output: name_emma_okonkwo

	fmt.Print(snaketoCamel("name_emma_okonkwo"))
	//Expected Output: nameEmmaOkonkwo
}
