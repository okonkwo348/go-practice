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
		return
	}

	if input[1] == "" {
		return
	}

	rows := strings.Split(input[1], "\\n")

	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	var banner []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		banner = append(banner, scanner.Text())
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
				start := int(char-32)*9 + 1
				fmt.Print(banner[start+i])
			}
			fmt.Println()
		}
	}

}
