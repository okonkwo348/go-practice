package main

import (
	"errors"
	"os"
	"strings"
)

func GenerateArt(text string, filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", errors.New("File Not Found")
	}

	splitLine := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	bannerLines := splitLine[1:]

	text = strings.ReplaceAll(text, "\r\n", "\n")
	sliceText := strings.Split(text, "\n")

	var b strings.Builder
	var result []string
	for _, word := range sliceText {
		for i := 0; i < 8; i++ {
			for _, ch := range word {
				pos := int(ch-32)*9 + i
				b.WriteString(bannerLines[pos])
			}
			result = append(result, b.String())
			b.Reset()
		}
	}
	return strings.Join(result, "\n"), nil
}

// func main() {
// 	art, _ := GenerateArt("go\nthere", "standard.txt")
// 	fmt.Println(art)
// }
