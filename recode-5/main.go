package main

import (
	"bufio"
	"fmt"
	"os"
)

func RenderAll(rows []string, banner map[rune][]string) {
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
				lines := banner[char]
				fmt.Print(lines[i])
			}
			fmt.Println()
		}
	}

}

func main() {
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("couldn't open file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var container []string
	for scanner.Scan() {
		container = append(container, scanner.Text())
	}

	banner := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		char := rune(i + 32)
		start := i*9 + 1
		lines := container[start : start+8]
		copyLines := make([]string, 8)
		copy(copyLines, lines)
		banner[char] = copyLines
	}

	RenderAll([]string{"hello", ""}, banner)

}
