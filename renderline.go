package main

import "fmt"

func RenderLine(s []string, substr string, bannerLine []string) {
	for i, text := range s {
		if text == "" {
			if i == len(s)-1 {
				continue
			}
			fmt.Println()
			continue
		}
		positions := findAllPositions(text, substr)
		for i := 0; i < 8; i++ {
			for charPos, char := range text {
				if isColored(charPos, positions, len(substr)) {
					lines := getCharLine(bannerLine, char)
					fmt.Print("\033[32m" + lines[i] + "\033[0m")
				} else {
					lines := getCharLine(bannerLine, char)
					fmt.Print(lines[i])
				}

			}
			fmt.Println()

		}

	}
}
