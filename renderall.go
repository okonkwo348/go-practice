package main

import "fmt"

func renderAll(rows []string, bannerLines []string) {
	for i, row := range rows {
		if row == "" {
			if i == len(row) {
				continue
			}
			fmt.Println()
			continue

		}

		fmt.Print(renderWord(row, bannerLines))
	}

}
