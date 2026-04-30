package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("couldn't load file")
	}

	if len(file) == 0 {
		return nil, errors.New("empty file")
	}

	if len(file) < 855 {
		return nil, errors.New(" lines out of range")
	}

	splitData := strings.Split(string(file), "\n")

	mapLine := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		char := rune(i + 32)
		start := i*9 + 1
		lines := splitData[start : start+8]

		copyLines := make([]string, 8)
		copy(copyLines, lines)

		mapLine[char] = copyLines
	}
	return mapLine, nil
}
