package main

import (
	"fmt"
	"strings"
)

func widthWords(gap, remain int) []int {
	space := []int{}

	if gap == 0 {
		return space
	}

	baseSpace := remain / gap
	extraSpace := remain % gap

	for i := 0; i < gap; i++ {
		space = append(space, baseSpace)
	}

	for j := 0; j < extraSpace; j++ {
		space[j] += 1
	}

	return space
}

func wordWidth(word string, banner map[rune][]string) int {
	widthOfWord := 0
	for _, char := range word {
		line := banner[char][0]
		widthOfWord += len(line)

	}

	return widthOfWord
}

func totalWordsWidth(words []string, banner map[rune][]string) int {
	totalWidthOfWords := 0

	for _, word := range words {
		totalWidthOfWords += wordWidth(word, banner)

	}

	return totalWidthOfWords
}

func justify(text string, banner map[rune][]string) {

	slice := strings.Fields(text)

	if len(slice) == 1 {
		for i := 0; i < 8; i++ {
			for _, char := range text {
				fmt.Print(banner[char][i])
			}
			fmt.Println()
		}
	}

	remainingSpace := 80 - totalWordsWidth(slice, banner)

	gap := len(slice) - 1

	positions := widthWords(gap, remainingSpace)

	for i := 0; i < 8; i++ {
		for _, word := range slice {

			for _, char := range text {
				fmt.Print(banner[char][i])
			}
			fmt.Println()
		}
	}

}
