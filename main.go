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
	flag = strings.TrimPrefix(flag, "--color=")

	mainString := os.Args[len(os.Args)-1]

	if len(os.Args) == 4 {
		subString := os.Args[2]
		bannerLine, _ := LoadBanner("standard.txt")
		valid, _ := ValidateInput(mainString)

		RenderLine(valid, subString, bannerLine)
	}

}
