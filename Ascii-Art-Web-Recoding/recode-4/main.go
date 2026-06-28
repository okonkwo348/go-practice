package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetCharLines(banner map[rune][]string, char rune) []string {
	result := make([]string, 8)
	copy(result, banner[char])
	return result

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

	fmt.Println(GetCharLines(banner, 'A'))
	fmt.Println(GetCharLines(banner, ' '))
	fmt.Println(GetCharLines(banner, '~'))
}
