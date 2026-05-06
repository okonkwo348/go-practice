package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Index(s, substr string) int{
	return strings.Index(s,substr)
}

func ValidateInput(s string) ([]string, error) {
	rows := strings.Split(s, "\\n")
	for _, word := range rows {
		for _, ch := range word {
			if ch < 32 || ch > 126 || ch == '\n' {
				return nil, errors.New("ch is not an ascii char")
			}
		}
	}
	return rows, nil
}

func LoadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	dataSplit := strings.Split(string(data), "\n")

	return dataSplit, nil
}

func getCharLine(bannerLine []string, char rune) []string {
	start := int(char-32)*9 + 1
	lines := bannerLine[start : start+8]
	charLine := make([]string, 8)
	copy(charLine, lines)
	return charLine
}

func RenderLine(s []string, sub string, bannerLine []string) {
	for i, text := range s {
		if text == "" {
			if i == len(s)-1 {
				continue
			}
			fmt.Println()
			continue
		}
			for i := 0; i < 8; i++ {
				if Index(text,sub){
					for _, char := range text {
					
						lines := getCharLine(bannerLine, char)
						fmt.Print("\033[32m" + lines[i] + "\033[0m")
				}
				
					} else {
						lines := getCharLine(bannerLine, char)
						fmt.Print(lines[i])
					}

				}
				fmt.Println()

			}

		}

	}
}

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("incomplete argument")
		return
	}

	flag := os.Args[1]
	flag = strings.TrimPrefix(flag, "--color=")

	mainString := os.Args[len(os.Args)-1]

	if len(os.Args) == 4 {
		subString := os.Args[2]
		bannerLine, _ := LoadBanner("standard.txt")
		valid, _ := ValidateInput(mainString)

		RenderLine(valid, subString, bannerLine)
	}

}
