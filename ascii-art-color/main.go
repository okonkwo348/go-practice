package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("incomplete argument")
		return
	}

	flag := os.Args[1]

	if !strings.HasPrefix(os.Args[1], "--color=") {
		fmt.Println("Usage: go run . --color=<color> <substring> \"string\"")
		return
	}

	flag = strings.TrimPrefix(flag, "--color=")

	colorCode := getColor(flag)

	if colorCode == "" {
		fmt.Println("invalide color - choose: red, green, blue, yellow, cryan,magenta, white")
		return
	}

	mainString := os.Args[len(os.Args)-1]

	subString := mainString

	if len(os.Args) == 4 {

		subString = os.Args[2]

	}

	bannerLine, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	valid, err1 := ValidateInput(mainString)
	if err1 != nil {
		fmt.Println("error: ", err1)
		return
	}

	RenderLine(valid, subString, colorCode, bannerLine)

}
