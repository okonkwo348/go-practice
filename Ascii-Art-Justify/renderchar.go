package main

func renderChar(bannerLine []string, char rune) []string {
	startPosition := (char-32)*9 + 1
	endPosition := startPosition + 8
	window := bannerLine[startPosition:endPosition]

	character := make([]string, 8)
	copy(character, window)

	return character

}
