package main

func widthWords(gap, remain int) []int {
	space := []int{}

	if gap == 0 {
		return space
	}

	basePadding := remain / gap
	extraSpace := remain % gap

	for i := 0; i < gap; i++ {
		space = append(space, basePadding)
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
