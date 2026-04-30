package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args

	if len(input) < 2 || len(input) > 3 {
		fmt.Println("argument not complete")
		return
	}

	if input[1] == "" {
		return
	}

	bannerFile := "standard.txt"

	if len(input) == 3 {
		validFileBanner := map[string]bool{
			"standard.txt":   true,
			"shadow.txt":     true,
			"thinkertoy.txt": true,
		}

		if !validFileBanner[input[2]] {
			fmt.Println("usage: choose from any of these;")
			return
		}

		bannerFile = input[2]
	}

	Maplines, _ := LoadBanner(bannerFile)

	lines := GenerateArt(input[1], Maplines)
	fmt.Println(lines)

}
