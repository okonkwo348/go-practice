package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("not complete arg")
		return
	}

	bannerFile := "standard" + ".txt"

	if len(os.Args) == 3 {
		mapBanner := map[string]bool{
			"standard":   true,
			"shadow":     true,
			"thinkertoy": true,
		}

		if !mapBanner[os.Args[2]] {
			fmt.Println("not among the available banner file. Select from standard, shadow and thinkertoy")
			return
		}

		bannerFile = os.Args[2] + ".txt"
	}

	_, err := ValidateInput(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	banner, err := LoadBanner(bannerFile)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	word := GenerateArt(os.Args[1], banner)
	fmt.Print(word)

}
