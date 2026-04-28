package main

import (
	"bufio"
	"errors"
	"os"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("file not found")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var container []string
	for scanner.Scan() {
		container = append(container, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.New("data not found")
	}

	if len(container) == 0 {
		return nil, errors.New("file is empty")
	}

	if len(container) < 855 {
		return nil, errors.New("invalid format")
	}

	mapLine := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		char := rune(i + 32)
		start := i*9 + 1

		lines := container[start : start+8]
		copyLines := make([]string, 8)
		copy(copyLines, lines)
		mapLine[char] = copyLines
	}

	return mapLine, nil
}
