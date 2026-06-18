package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error in reading file")
	}

	bannerLines := strings.Split(string(data), "\n")
    
    if len(bannerLines)==0{
        return nil, errors.New("Empty File")
    }

    if len(bannerLines)<855{
        return nil, errors.New("incomplete lines")
    }
    
	var lines []string
	mapLine := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		char := rune(i + 32)
		start := i*9 + 1
		lines = bannerLines[start : start+8]

		mapLine[char] = lines

	}

	return mapLine, nil
}
