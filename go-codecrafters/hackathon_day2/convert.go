package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Hex -> Dec
func hex(s string) string {
	a, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return "invalid input"
	}
	return strconv.FormatInt(a, 10)
}

// Bin -> Dec
func bin(s string) string {
	a, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return "invalid input"
	}
	return strconv.FormatInt(a, 10)

}

// Dec -> Hex
func decHex(s int64) string {

	return strconv.FormatInt(s, 16)
}

func decBin(s int64) string {

	return strconv.FormatInt(s, 2)
}

func main() {
	fmt.Println("WELCOME CONVERSION CALCULATOE")
	fmt.Println()
	fmt.Println("U")
	fmt.Println("Hex -> Dec do : convert 1E hex")
	fmt.Println("Bin -> Dec do : convert 10 bin")
	fmt.Println("Dec -> Hex or Bin do : convert 1E dec")
start:
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		goto start
	}

	line := scanner.Text()
	part := strings.Fields(line)

	if len(part) != 3 && part[0] != "quit" && part[0] != "HELP" && part[0] != "help" {
		fmt.Println("error: requires three arguments, please type [HELP] for clarity")
		goto start
	}

	if part[0] == "quit" {
		fmt.Println("exit")
	}

	if part[0] == "HELP" || part[0] == "help" {
		fmt.Println("Hex -> Dec do : convert 1E hex")
		fmt.Println("Bin -> Dec do : convert 10 bin")
		fmt.Println("Dec -> Hex or Bin do : convert 1E dec")
		fmt.Println()
		goto start
	}

	if len(part) == 3 && part[0] == "convert" && strings.ContainsAny(part[1], "0123456789ABCDEF") || part[2] == "hex" || part[2] == "bin" || part[2] == "dec" {
		switch part[2] {
		case "bin":
			fmt.Println(bin(part[1]))
			fmt.Println()
			goto start
		case "hex":
			fmt.Println(hex(part[1]))
			fmt.Println()
			goto start
		case "dex":
			fmt.Println(part[1])
			fmt.Println()
			goto start
		default:
			fmt.Println("out the given functions")
			fmt.Println()
			goto start
		}
	}
}
