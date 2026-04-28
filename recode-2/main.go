package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ValidateInput(input string) ([]string, error) {
	splitChar := strings.Split(input, "\\n")

	for _, row := range splitChar {
		for _, char := range row {
			if !(char >= 32 && char <= 126) {
				return nil, errors.New("input not in ascii table ")
			}

		}
	}
	return splitChar, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go \"text\"")
		return
	}

	input := os.Args[1]
	valid, err := ValidateInput(input)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("%#v", valid)

}
