package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args

	if len(input) != 2 {
		fmt.Println("arguments command incomplete")
		return
	}

	if input[1] == "" {
		return
	}

	rows := strings.Split(input[1], "\\n")

	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("couldn't open the file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bannerLine []string
	for scanner.Scan() {
		bannerLine = append(bannerLine, scanner.Text())
	}

	linesMap := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		char := rune(i + 32)
		start := i*9 + 1
		window := bannerLine[start : start+8]
		lines := make([]string, 8)
		copy(lines, window)

		linesMap[char] = lines
	}

	for i, row := range rows {
		if row == "" {
			if i == len(rows)-1 {
				continue
			}
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range row {
				fmt.Print(linesMap[char][i])

			}
			fmt.Println()

		}
	}

}
