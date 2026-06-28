package main

import (
	"strings"
)

func renderAll(rows []string, bannerLines []string) string {
	var result strings.Builder
	for i, row := range rows {
		if row == "" {
			if i == len(rows)-1 {
				continue
			}
			result.WriteString("\n")
			continue

		}
		result.WriteString(renderWord(row, bannerLines))

	}

	return result.String()

}
