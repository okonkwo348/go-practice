package main

import (
	// "bufio"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func validate(s []string) ([]string, error) {
	for _, row := range s {
		for _, val := range row {
			if !(val >= 32 && val <= 126 || val == '\n') {
				return nil, errors.New("invalid character not a ascii char")
			}
		}
	}

	return s, nil
}

func fileHandling(banner string) ([]string, error) {
	file, err := os.Open(banner)
	if err != nil {
		return nil, errors.New("couldn't open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	container := []string{}
	for scanner.Scan() {
		container = append(container, scanner.Text())
	}

	return container, nil
}

func getEachChar(banner []string, char rune) []string {
	startPosition := (char-32)*9 + 1
	endPosition := startPosition + 8

	actulPosition := banner[startPosition:endPosition]

	return actulPosition
}

func eachWord(s string, banner []string) string {
	var result strings.Builder
	for i := 0; i < 8; i++ {
		for _, row := range s {
			line := getEachChar(banner, row)
			result.WriteString(line[i])
		}
		result.WriteString("\n")
	}
	return result.String()
}
func allWords(rows []string, banner []string) {
	for i, row := range rows {
		if row == "" {
			if i == len(rows)-1 {
				continue
			}
			fmt.Println()
			continue
		}
		fmt.Println(eachWord(row, banner))
	}

}
func main() {
	input := os.Args

	if len(input) != 2 {
		fmt.Println("incomplete commandline")
		return
	}

	if input[1] == "" {
		return
	}

	splitChar := strings.Split(input[1], "\\n")
	net, err := validate(splitChar)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	now, err1 := (fileHandling("banner/standard.txt"))
	if err1 != nil {
		fmt.Println("Error:", err1)
		return
	}
	allWords(net, now)
}
