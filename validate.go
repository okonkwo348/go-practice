package main

import "errors"

func ValidateInput(s string) (rune, error) {
	for _, ch := range s {
		if ch < 32 || ch > 126 {
			return ch, errors.New("ch is not an ascii char ")
		}
	}
	return 0, nil
}
