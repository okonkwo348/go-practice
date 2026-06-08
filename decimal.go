package main

import (
	"fmt"
	"strconv"
)

func hexToDec(s string) (int64, error) {
	data, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("%v is not acceptable", err)
	}

	return data, nil
}

func main() {
	data, err := hexToDec("BADFOOD")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(data)
}
