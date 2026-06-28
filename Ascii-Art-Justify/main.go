package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("incomplete argument. Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	mainString := ""

	if len(os.Args) == 2 {
		mainString = os.Args[1]
	}

	bannerStyle := "standard" + ".txt"

	flag := "left"

	if len(os.Args) == 4 || (len(os.Args) == 3 && strings.HasPrefix(os.Args[1], "--align=")) {

		mainString = os.Args[2]

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

	if len(os.Args) == 4 || (len(os.Args) == 3 && !strings.HasPrefix(os.Args[1], "--align=")) {
		ValidBanner := map[string]bool{
			"standard":   true,
			"shadow":     true,
			"thinkertoy": true,
		}

		if len(os.Args) == 3 {
			if !ValidBanner[os.Args[2]] {
				fmt.Println("not among the available banner file. Select from standard, shadow and thinkertoy")
				return
			}
			bannerStyle = os.Args[2] + ".txt"
			mainString = os.Args[1]
		}

		if len(os.Args) == 4 {
			if !ValidBanner[os.Args[3]] {
				fmt.Println("not among the available banner file. Select from standard, shadow and thinkertoy")
				return
			}
			bannerStyle = os.Args[3] + ".txt"
		}

	}

	// 1. load banner
	banner, err := loadBanner("banner/" + bannerStyle)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// 2. get terminal width
	width := getTerminalWidth()

	// 3 & 4. render and align each line
	lines := strings.Split(mainString, "\\n")

	validLines, err := validateInput(lines)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var finalOutput strings.Builder

	for _, line := range validLines {
		var blocks []string
		if flag == "justify" {
			blocks = renderWords(line, banner)
		} else {
			blocks = []string{RenderWord(line, banner)}
		}

		finalOutput.WriteString(Alignment(blocks, width, flag))
	}

	fmt.Print(finalOutput.String())
}
