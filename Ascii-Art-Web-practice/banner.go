package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}

	dataFile := strings.Split(string(data), "\n")

	if len(dataFile) == 0 {
		return nil, errors.New("Empty file")
	}

	if len(dataFile) < 855 {
		return nil, errors.New("File is currupted")
	}

	return dataFile, nil

}
