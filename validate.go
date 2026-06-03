package main

import (
	"fmt"
)

func ValidateInput(input []string) ([]string, error) {
	for _, word := range input {
		for _, char := range word {
			if (char < 32 || char > 126) && char != '\n' {
				return nil, fmt.Errorf("%v not an ascii character", char)
			}
		}
	}
	return input, nil
}
