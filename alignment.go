package main

import "strings"

func Alignment(renderArt []string, terminalWidth int, flag string) string {
	switch flag {
	case "left":
		var result strings.Builder

		// first, split each block into rows
		var allLines [][]string
		for _, block := range renderArt {
			lines := strings.Split(block, "\n")
			allLines = append(allLines, lines)
		}

		// then combine row by row
		for i := 0; i < 8; i++ {
			for _, charLines := range allLines {
				result.WriteString(charLines[i])
			}
			result.WriteString("\n")
		}
	case "right":
		var result strings.Builder

		var allLines [][]string
		for _, block := range renderArt {
			lines := strings.Split(block, "\n")
			allLines = append(allLines, lines)
		}

		// calculate total text width first

		totalWidth := 0
		for _, lines := range allLines {
			totalWidth += len(lines[0])
		}

		padding := terminalWidth - totalWidth
		leftPad := strings.Repeat(" ", padding)

		for i := 0; i < 8; i++ {
			result.WriteString(leftPad)
			for _, charLines := range allLines {
				result.WriteString(charLines[i])
			}
			result.WriteString("\n")
		}

	case "center":
		var result strings.Builder

		var allLines [][]string
		for _, block := range renderArt {
			lines := strings.Split(block, "\n")
			allLines = append(allLines, lines)
		}

		// calculate toal text width first
		totalWidth := 0
		for _, lines := range allLines {
			totalWidth += len(lines[0])
		}
		padding := terminalWidth - totalWidth

		leftPad := strings.Repeat(" ", padding/2)
		rightPad := strings.Repeat(" ", padding-padding/2)

		for i := 0; i < 8; i++ {
			result.WriteString(leftPad)
			for _, charLines := range allLines {
				result.WriteString(charLines[i])

			}
			result.WriteString(rightPad)
			result.WriteString("\n")

		}
	case "justify":
		var result strings.Builder

		var allLines [][]string
		for _, block := range renderArt {
			lines := strings.Split(block, "\n")
			allLines = append(allLines, lines)
		}

		totalWidth := 0
		for _, lines := range allLines {
			totalWidth += len(lines[0])
		}

		totalSpace := terminalWidth - totalWidth
		gaps := len(allLines) - 1
		spacePerGap := totalSpace / gaps
		leftover := totalSpace % gaps

		for i := 0; i < 8; i++ {
			if gaps == 0 {
				for _, charLines := range allLines {
					result.WriteString(charLines[i])
				}
			} else {

				for j, charLines := range allLines {
					result.WriteString(charLines[i])
					if j < gaps {
						extra := 0
						if j < leftover {
							extra = 1
							result.WriteString(strings.Repeat(" ", spacePerGap+extra))
						}
					}
				}
			}
			result.WriteString("\n")
		}

	default:
		return "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard\n"

	}

}
