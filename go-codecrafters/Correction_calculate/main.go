package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	return a / b
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter an operation from the below")
	fmt.Println()
	fmt.Println("TO ADD do -> add 4667 544")
	fmt.Println("TO SUBTRATE do -> sub 467 54")
	fmt.Println("TO MULTIPLY do -> mul 4667 544")
	fmt.Println("TO DIVIDE do -> div 46 4")
	fmt.Println("TO STOP THE PROGRAM type: quit")
	fmt.Println("NEED HELP type: help")
	fmt.Println()

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		part := strings.Fields(input)

		if len(part) == 0 {
			fmt.Println("empty: must enter three elemnets")
			continue
		}

		if part[0] == "help" {
			fmt.Println("Enter an operation from the below")
			fmt.Println()
			fmt.Println("TO ADD type -> add 4667 544")
			fmt.Println("TO SUBTRATE type -> sub 467 54")
			fmt.Println("TO MULTIPLY type -> mul 4667 544")
			fmt.Println("TO DIVIDE type -> div 46 4")
			fmt.Println("TO STOP THE PROGRAM type: quit")
			fmt.Println("NEED HELP type: help")
			fmt.Println()
			continue
		}

		if part[0] == "quit" {
			break
		}

		// hint — check against valid operations
		if part[0] != "add" && part[0] != "sub" && part[0] != "mul" && part[0] != "div" {
			fmt.Println("unknown command — type help for usage")
			continue
		}
		if len(part) != 3 {
			fmt.Println("elemets incomplete pls must enter three elemnets")
			continue
		}

		a, err := strconv.ParseFloat(part[1], 64)
		if err != nil {
			fmt.Println("first value must be a number")
			continue

		}

		b, err2 := strconv.ParseFloat(part[2], 64)
		if err2 != nil {
			fmt.Println("second value must be a number")
			continue

		}
		switch part[0] {
		case "add":
			fmt.Printf("Result: %.2f\n", add(a, b))
			continue

		case "sub":
			fmt.Printf("Result: %.2f\n", sub(a, b))
			continue

		case "mul":
			fmt.Printf("Result: %.2f\n", mul(a, b))
			continue

		case "div":
			if b == 0 {
				fmt.Println("Can't divide by zero")
				continue
			}

			fmt.Printf("Result: %.2f\n", div(a, b))
			continue

		default:
			fmt.Println("invalid")
			continue

		}
	}
}
