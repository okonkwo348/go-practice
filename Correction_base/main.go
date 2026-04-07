package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bin(s string) (string, error) {
	a, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(a, 10), nil

}

func hex(s string) (string, error) {
	a, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(a, 10), nil

}

func decHex(s string) (string, error) {
	a, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(strconv.FormatInt(a, 16)), nil
}

func decBin(s string) (string, error) {
	a, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(a, 2), nil
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

		if len(part) != 3 {
			fmt.Println("elemets incomplete pls must enter three elemnets")
			continue
		}

		if part[0] != "convert" {
			fmt.Println("invalid first word")
			continue
		}

		if part[2] != "hex" && part[2] != "bin" && part[2] != "dec" {
			fmt.Println("invalid last char")
			continue
		}

		switch part[2] {
		case "bin":
			nib, err := bin(part[1])
			if err != nil {
				fmt.Println("invalid second char for bin")
				continue
			}
			fmt.Printf("Decimal: %s\n", nib)
			continue
		case "hex":
			xeh, err1 := hex(part[1])
			if err1 != nil {
				fmt.Println("invalid second char for hex")
				continue
			}
			fmt.Printf("Decimal: %s\n", xeh)
			continue

		case "dec":
			ced, err2 := decBin(part[1])
			if err2 != nil {
				fmt.Println("invalid second char for dec")
				continue
			}

			ced1, err3 := decHex(part[1])
			if err3 != nil {
				fmt.Println("invalid second char for dec")
				continue
			}

			fmt.Printf("Binary: %s\n", ced)
			fmt.Printf("Hex: %s\n", ced1)
			continue
		default:
			fmt.Printf("Invalid base\n")
		}

	}
}
