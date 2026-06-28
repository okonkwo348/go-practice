package main

import (
	"fmt"
	"strings"
)

func ValidateInput(s string) (string, error) {
	input := strings.ReplaceAll(s, "\r\n", "")
	for _, char := range input {
		if char < 32 || char > 126 {
			return "", fmt.Errorf("%v not an ascii character", char)
		}
	}
	return s, nil
}
