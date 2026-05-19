package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("incomplete argument")
		return
	}

	mainString := ""

	if len(os.Args) == 2 {
		mainString = os.Args[1]
	}

	bannerStyle := "standard" + ".txt"

	flag := "left"

	if len(os.Args) > 2 {
		flag = strings.TrimPrefix(os.Args[1], "--align=")

		validFlags := map[string]bool{
			"left":    true,
			"right":   true,
			"center":  true,
			"justify": true,
		}

		if !validFlags[flag] {
			fmt.Println("Usage: go run . --align=<left|right|center|justify> \"string\" [banner]")
			return
		}
	}

	if len(os.Args) == 3 {
		mainString = os.Args[2]
	}

	if len(os.Args) == 4 {
		ValidBanner := map[string]bool{
			"standard":   true,
			"shadow":     true,
			"thinkertoy": true,
		}

		if !ValidBanner[os.Args[3]] {
			fmt.Println("not among the available banner file. Select from standard, shadow and thinkertoy")
			return
		}

		bannerStyle = os.Args[3] + ".txt"

		mainString = os.Args[2]

	}

}
