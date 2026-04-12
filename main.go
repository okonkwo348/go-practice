package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func validateInput(arg []string) ([]string, error) {

	for _, row := range arg {
		for _, char := range row {
			if !(char >= 32 && char <= 126 || char == '\n') {
				return nil, errors.New("invalid input")
			}

		}
	}
	return arg, nil
}

func main() {
	input := os.Args

	// Chang if argument is provided
	if len(input) != 2 {
		return
	}

	splitInput := strings.Split(input[1], "\\n")

	valid, err := validateInput(splitInput)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	bannerLine, err := loadBanner("banner/standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	renderAll(valid, bannerLine)

}
