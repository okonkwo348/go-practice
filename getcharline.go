package main

func getCharLine(bannerLine []string, char rune) []string {
	start := int(char-32)*9 + 1
	lines := bannerLine[start : start+8]
	charLine := make([]string, 8)
	copy(charLine, lines)
	return charLine
}
