package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ValidateInput(s string) ([]string, error) {
	rows := strings.Split(s, "\\n")
	for _, word := range rows {
		for _, ch := range word {
			if ch < 32 || ch > 126 {
				return nil, errors.New("char is not an ascii character")
			}
		}
	}
	return rows, nil
}

func main() {
	input := os.Args

	if len(input) < 2 || len(input) > 3 {
		fmt.Println("usage: go run \"string\" [standaer][shadow][thinkertoy]")
		return
	}

	if input[1] == "" {
		return
	}

	rows, _ := ValidateInput(input[1])

	bannerFile := "standard.txt"

	if len(input) == 3 {
		validateBannerFile := map[string]bool{
			"standard":   true,
			"shadow":     true,
			"thinkertoy": true,
		}

		if !validateBannerFile[input[2]] {
			fmt.Println("invalid file: choose either shadow or standard or thinkertoy")
			return
		}

		bannerFile = input[2] + ".txt"
	}

	data, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("couldn't read banner file")
		return
	}

	if len(data) == 0 {
		fmt.Println("empty file")
		return
	}

	if len(data) < 855 {
		fmt.Println("lines of slice out of bound")
		return
	}

	datas := strings.ReplaceAll(string(data), "\r\n", "\n")

	splitLines := strings.Split(datas, "\n")

	for i, row := range rows {
		if row == "" {
			if i == len(rows)-1 {
				continue
			}
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, c := range row {
				start := (int(c)-32)*9 + 1
				fmt.Print(splitLines[start+i])
			}
			fmt.Println()
		}
	}

}
