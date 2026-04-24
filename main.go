package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	input := os.Args

	if len(input) != 2 {
		fmt.Println("Argument is not com[plete]")
		return
	}

	if input[1] == "" {
		fmt.Println()
		return
	}

	rows := strings.Split(input[1], "\\n")

	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("couldn't open banner file")
		return
	}

	banner := strings.Split(string(file), "\n")
	mapArt := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		start := i*9 + 1
		char := rune(i + 32)
		mapArt[char] = banner[start : start+8]
	}

	for i, row := range rows {
		if row == "" {
			if i == len(row)-1 {
				continue
			}
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range row {
				fmt.Print(mapArt[char][i])
			}
			fmt.Println()
		}
	}
}
