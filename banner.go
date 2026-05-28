package main

import (
	"bufio"
	"errors"
	"os"
)

func loadBanner(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("error at openning the file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bannerLine []string
	for scanner.Scan() {
		bannerLine = append(bannerLine, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("error reading file content")
	}

	return bannerLine, nil
}
