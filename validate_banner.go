package main

import (
	"errors"
	"fmt"
)

func ValidateBanner(banner map[rune][]string) error {
	if banner == nil {
		return errors.New("banner is new")
	}

	if len(banner) != 95 {
		return fmt.Errorf("banner has %d entries, expected 95", len(banner))
	}

	for r := rune(32); r <= 126; r++ {
		val, ok := banner[r]

		if !ok {
			return fmt.Errorf("missing character: %c", r)
		}

		if len(val) != 8 {
			return fmt.Errorf("character 'A' has %d lines, expected 8", len(val))
		}
	}
	return nil
}
