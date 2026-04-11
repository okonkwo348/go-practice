package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadBanner(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)

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

	character := bannerLine[startPosition:endPosition]

	result := []string{}
	for i := 0; i < 8; i++ {
		for _, char := range character {
			result = append(result, char)
		}
		result = append(result, "\n")
	}
	return result

}
