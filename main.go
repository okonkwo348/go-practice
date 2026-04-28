package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//mapLine, _ := LoadBanner("standard.txt")
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Error: ")
		return
	}

	row := strings.Split(os.Args[1], `\n`)
	valid, err := ValidateInput(row)
	if err != nil {
		fmt.Println("Error: ")
		return
	}

	banner, _ := LoadBanner("standard.txt")
	for _, v := range valid {
		fmt.Print(v, banner)
	}

	// for i := 32; i <= 126; i++ {
	// 	fmt.Printf("%c\n", i)
	// 	fmt.Println()
	// 	fmt.Printf("%s", strings.Join(mapLine[rune(i)], "\n"))
	// }

	//fmt.Print(RenderLine("k", mapLine))
}
