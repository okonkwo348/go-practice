package main

import (
	"errors"
	"fmt"
	"os"
)

func ReadFile(name string) (string, error) {

	file, err := os.ReadFile(name)
	if err != nil {
		return "", errors.New("file Not Found")
	}

	return string(file), nil
}

func main() {
	input := os.Args

	if input[1] == "" {
		println("No argument is provided")
	}

	data, err := ReadFile(input[1])
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Println(data)

}
