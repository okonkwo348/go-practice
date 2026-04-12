package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func loadBanner(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, errors.New("error at openning the file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bannerLine []string
	for scanner.Scan() {
		bannerLine = append(bannerLine, scanner.Text())
	}

	return bannerLine, nil
}

func getCharLines(bannerLine []string, char rune) []string {
	startPosition := (char-32)*9 + 1
	endPosition := startPosition + 8
	window := bannerLine[startPosition:endPosition]

	character := make([]string, 8)
	copy(character, window)

	return character

}

func renderWord(word string, bannerLines []string) string {
	var result strings.Builder
	for i := 0; i < 8; i++ {
		for _, char := range word {
			line := getCharLines(bannerLines, char)
			result.WriteString(line[i])
		}
		result.WriteString("\n")
	}

	return result.String()
}

func renderAll(rows []string, bannerLines []string) {
	for _, row := range rows {
		if row == "" {
			fmt.Println()
			continue
		}
		fmt.Print(renderWord(row, bannerLines))
	}

}
