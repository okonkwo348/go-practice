package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(file string) (map[string]string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.New("file not found")
	}

	splitDate := strings.Split(string(data), "\n")

	mapLine := make(map[string]string)
	var lines string

	for i := 'A'; i <= 'C'; i++ {
		char := i
		line := i - 'A'
		lines = splitDate[line]
		mapLine[string(char)] = lines
	}

	return mapLine, nil
}
