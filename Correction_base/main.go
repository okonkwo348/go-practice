package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bin(s string) string {
	a, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return strconv.FormatInt(a, 10)

}

func hex(s string) string {
	a, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return strconv.FormatInt(a, 10)

}

func decHex(s string) string {
	a, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return strconv.FormatInt(a, 16)
}

func decBin(s string) string {
	a, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return strconv.FormatInt(a, 2)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		part := strings.Fields(input)

		if len(part) == 0 {
			fmt.Println("incomplete characters")
			continue
		}

		if part[0] == "quit" {
			break
		}

		if part[0] != "convert" {
			fmt.Println("invalid first word")
		}

		if part[1] == "hex" && !strings.ContainsAny(part[1], "abcdefABCDEF") {
			fmt.Println("invalid for Hex")
			continue
		}

		if part[1] == "bin" && !strings.ContainsAny(part[1], "01") {
			fmt.Println("invalid for bin")
			continue
		}

		if part[1] == "dec" && !strings.ContainsAny(part[1], "abc") {
			fmt.Println("invalid for dec")
			continue
		}

		switch part[2] {
		case "bin":
			fmt.Printf("Decimal: %s", bin("10"))
			continue
		case "hex":
			fmt.Printf("Decimal: %s", bin("1E"))
			continue

		case "dec":
			fmt.Printf("Binary: %s", decBin("255"))
			fmt.Printf("Hex: %s", decHex("255"))
			continue
		default:
			fmt.Printf("Invalid base")
		}

	}
}
