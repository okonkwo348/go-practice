package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a, b string) string {
	c, err1 := strconv.Atoi(a)
	d, err2 := strconv.Atoi(b)
	e := c + d
	if err1 != nil {
		return "wrong first value"
	}
	if err2 != nil {
		return "wrong second value"
	}
	return strconv.Itoa(e)

}
func sub(a, b string) string {
	c, err1 := strconv.Atoi(a)
	d, err2 := strconv.Atoi(b)
	e := c - d
	if err1 != nil {
		return "wrong first value"
	}
	if err2 != nil {
		return "wrong second value"
	}
	return strconv.Itoa(e)

}

func mul(a, b string) string {
	c, err1 := strconv.Atoi(a)
	d, err2 := strconv.Atoi(b)

	e := c * d
	if err1 != nil {
		return "wrong first value"
	}
	if err2 != nil {
		return "wrong second value"
	}
	if c == 0 || d == 0 {
		return "0"
	}
	return strconv.Itoa(e)
}

func div(a, b string) string {
	c, err1 := strconv.Atoi(a)
	d, err2 := strconv.Atoi(b)

	e := c / d
	if err1 != nil {
		return "wrong first value"
	}
	if err2 != nil {
		return "wrong second value"
	}
	if d == 0 {
		return "can't"
	}
	if c == 0 {
		return "0"
	}

	return strconv.Itoa(e)
}

func main() {
	fmt.Println("Welcome to the master calculator")
	fmt.Println()
	fmt.Println("follow the pattern from the give examples below")

	fmt.Println("To ADD do: <sum> <num1> <num2>   e.g add 3 4")
	fmt.Println("To SUBSTRATE do: <sub> <num1> <num2>  e.g sub 53 79")
	fmt.Println("To MULTIPLY do: <mul> <num1> <num2>    e.g mul 34 334")
	fmt.Println("To DIVIDE do: <div> <num1> <num2>   e.g div 333 455")
	fmt.Println()
	fmt.Println("To QUIT do: quit")
	fmt.Println("need help do: HELP")
	fmt.Println()

start:
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("start by type Here:")
	if !scanner.Scan() {
		goto start
	}

	line := scanner.Text()
	part := strings.Fields(line)

	if part[0] == "quit" {
		fmt.Println("exit")
	}

	if part[0] == "HELP" || part[0] == "help" {
		fmt.Println("To ADD do: <sum> <num1> <num2>   e.g add 3 4")
		fmt.Println("To SUBSTRATE do: <sub> <num1> <num2>  e.g sub 53 79")
		fmt.Println("To MULTIPLY do: <mul> <num1> <num2>    e.g mul 34 334")
		fmt.Println("To DIVIDE do: <div> <num1> <num2>   e.g div 333 455")
		fmt.Println()
		goto start
	}
	if len(part) != 3 && part[0] != "quit" && part[0] != "HELP" && part[0] != "help" {
		fmt.Println("error: requires three arguments, please type [HELP] for clarity")
		fmt.Println()
		goto start
	}

	if len(part) == 3 && part[0] == "add" || part[0] == "sub" || part[0] == "mul" || part[0] == "div" || part[0] == "add" {
		switch part[0] {
		case "add":
			fmt.Println(add(part[1], part[2]))
			fmt.Println()
			goto start
		case "sub":
			fmt.Println(sub(part[1], part[2]))
			fmt.Println()
			goto start
		case "mul":
			fmt.Println(mul(part[1], part[2]))
			fmt.Println()
			goto start
		case "div":
			if part[2] == "0" {
				fmt.Println("can't divide by zero")
				fmt.Println()
				goto start
			}
			fmt.Println(div(part[1], part[2]))
			fmt.Println()
			goto start
		default:
			fmt.Println("usage: <sum><sub><mul><div> <num1> <num2>")
			fmt.Println()
			goto start
		}
	} else if len(part) == 3 && part[0] != "quit" && part[0] != "HELP" && part[0] != "help" {
		fmt.Println("usage pattern: <sum or sub or mul or div> <num1> <num2>")
		fmt.Println()
		goto start
	}
}
