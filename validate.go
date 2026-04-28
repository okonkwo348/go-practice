package main 

import (
	"fmt"
)

func ValidateInput(input []string) ([]string, error) {
	for _, row := range input {
		for _, ch := range row {
			if !(ch >= 32) && !(ch <= 126) {
				return nil, fmt.Errorf("not a ascii character")
			}
		}
	}
	return input, nil 
}