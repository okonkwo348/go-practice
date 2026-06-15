package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args

	valid, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	for _, char := range input[1] {
		_, ok := valid[string(char)]
		if ok {
			fmt.Println(valid[string(char)])
		} else {
			fmt.Printf("%s not found", string(char))
		}
	}

}
