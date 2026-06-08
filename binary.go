package main

import (
	"fmt"
	"strconv"
)

func binToDec(s string) (int64, error) {
	return strconv.ParseInt(s, 2, 64)
}

func main() {
	data, err := binToDec("111111")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Println(data)
}
