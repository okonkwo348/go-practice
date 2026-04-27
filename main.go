package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RenderWord(word string, banner map[rune][]string) string {
	var result strings.Builder
	for i := 0; i < 8; i++ {
		for _, char := range word {
			result.WriteString(banner[char][i])
		}
		result.WriteString("\n")
	}
	return result.String()
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

	fmt.Println(RenderWord("Hi", banner))
	fmt.Println(RenderWord("", banner))
	fmt.Println(RenderWord("A B", banner))
}
