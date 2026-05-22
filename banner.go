package main

import (
	"bufio"
	"errors"
	"os"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("file doesn't exist or corrupt")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var container []string
	for scanner.Scan() {
		container = append(container, scanner.Text())
	}

	mapLine := make(map[rune][]string)
	var line []string

	for i := ' '; i < '~'; i++ {
		start := (int(i)-32)*9 + 1
		line = container[start : start+8]
		lines := make([]string, 8)
		copy(lines, line)
		mapLine[i] = lines
	}

	return mapLine, nil

}
