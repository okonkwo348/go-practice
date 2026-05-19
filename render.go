package main

import (
	"fmt"
	"strings"
)

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
	for i, row := range rows {
		if row == "" {
			if i == len(row) {
				continue
			}
			fmt.Println()
			continue

		}

	}

	fmt.Print(justify(rows, bannerLines))

}

func justify(input []string, bannerLine []string) {

	if len(input) == 1 {
		for i := 0; i < 8; i++ {
			for _, char := range input[0] {
				lines := getCharLines(bannerLine, char)
				fmt.Print(lines[i])
			}
			fmt.Println()
		}
		return
	}

	remainingSpace := 80 - totalWordsWidth(input, bannerLine)
	gap := len(input) - 1
	positions := widthWords(gap, remainingSpace)
	for i := 0; i < 8; i++ {
		for j, word := range input {
			for _, char := range word {
				lines := getCharLines(bannerLine, char)
				fmt.Print(lines[i])
			}
			if j != len(input)-1 {
				fmt.Print(strings.Repeat(" ", positions[j]))
			}
		}
		fmt.Println()
	}
}
