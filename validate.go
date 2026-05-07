package main

import (
	"errors"
	"strings"
)

func ValidateInput(s string) ([]string, error) {
	rows := strings.Split(s, "\\n")
	for _, word := range rows {
		for _, ch := range word {
			if ch < 32 || ch > 126 || ch == '\n' {
				return nil, errors.New("ch is not an ascii char")
			}
		}
	}
	return rows, nil
}
