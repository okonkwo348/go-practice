package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args

	if len(input) < 2 || len(input) > 4 {
		fmt.Println("argument not complete")
		return
	}

	mainString := ""

	if len(os.Args) == 2 {
		mainString = os.Args[1]
	}

	flag := "result.txt"

	if len(os.Args) == 4 || (len(os.Args) == 3 && strings.HasPrefix(os.Args[1], "--output=")) {
		flag = strings.TrimPrefix(os.Args[1], "--output=")
		mainString = os.Args[2]

	}

	bannerFile := "standard" + ".txt"

	if len(os.Args) == 4 || (len(os.Args) == 3 && !strings.HasPrefix(os.Args[1], "--output=")) {
		validFileBanner := map[string]bool{
			"standard":   true,
			"shadow":     true,
			"thinkertoy": true,
		}

		if len(os.Args) == 3 {
			if !validFileBanner[input[2]] {
				fmt.Println("usage: choose from any of these;")
				return
			}
			mainString = os.Args[1]
			bannerFile = input[2] + ".txt"
		}

		if len(os.Args) == 4 {
			if !validFileBanner[input[3]] {
				fmt.Println("usage: choose from any of these;")
				return
			}
			bannerFile = input[3] + ".txt"
		}

	}

	Maplines, _ := LoadBanner(bannerFile)

	lines := GenerateArt(mainString, Maplines)

	err2 := os.WriteFile(flag, []byte(lines), 0644)
	if err2 != nil {
		fmt.Println("file couldn't write")
		return
	}

}
