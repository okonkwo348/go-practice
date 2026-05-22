package main

import (
	"errors"
)

func validateInput(arg []string) ([]string, error) {

	for _, row := range arg {
		for _, char := range row {
			if !(char >= 32 && char <= 126 || char == '\n') {
				return nil, errors.New("invalid input not part of ascii character")
			}

		}
	}
	return arg, nil
}
