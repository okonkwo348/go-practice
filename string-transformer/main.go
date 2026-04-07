package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Upper(s string) string {
	return strings.ToUpper(s)
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func onWordCap(a string) string {
	s := strings.ToLower(a)
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func Cap(s string) string {
	a := strings.Fields(s)
	for i := 0; i < len(a); i++ {
		a[i] = onWordCap(a[i])
	}
	return strings.Join(a, " ")
}

func Title(a string) string {
	s := strings.Fields(a)
	for i := 0; i < len(s); i++ {
		if len(s[i]) > 3 || i == 0 {
			s[i] = onWordCap(s[i])
		}
	}
	return strings.Join(s, " ")
}

func revWord(a string) string {
	runes := []rune(a)
	result := []rune{}
	for i := len(runes) - 1; i >= 0; i-- {
		rev := runes[i]
		result = append(result, rev)
	}
	return string(result)
}

func Reverse(a string) string {
	s := strings.Fields(a)
	result := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		rev := revWord(s[i])
		result = append(result, rev)
	}
	return strings.Join(result, " ")
}

func Snake(a string) string {
	s := []rune(a)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			s[i] = '_'
		case ',':
			s[i] = '_'
		case '.':
			s[i] = '_'
		}
	}
	return string(s)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to string Transformation\n")
	fmt.Println("Follow the format below: \n")
	fmt.Println("<command> <text>")
	fmt.Println("for UPPERCASE type: upper <text> --->  upper sentinel is watching ")
	fmt.Println("For LOWERCASE type: lower <text> --->  ALERT LEVEL FIVE DETECTED")
	fmt.Println("For CAPITALIZATION type: cap <text>  ---->  cap director adaeze okonkwo")
	fmt.Println("For TILE type: tile <text>  ---->  title the rise of the eastern threat")
	fmt.Println("For SNAKE type: snake <text>  ---->  snake Operation Gopher Protocol!")
	fmt.Println("For REVERSE type: reverse <text>  ---->  reverse Go saves the world")

	for {
		fmt.Print("\n> ")
		scanner.Scan()
		input := scanner.Text()

		part := strings.Fields(input)

		if len(part) < 1 {
			fmt.Println("Empty: Type from one the commands")
			continue
		}

		if part[0] == "quit" {
			fmt.Println(" Shutting down String Transformer. Goodbye.")
			break
		}
		if len(part) < 2 && part[0] != "upper" && part[0] != "UPPER" && part[0] != "lower" && part[0] != "LOWER" &&
			part[0] != "cap" && part[0] != "CAP" && part[0] != "title" && part[0] != "TITLE" && part[0] != "snake" &&
			part[0] != "SNAKE" && part[0] != "reverse" && part[0] != "REVERSE" {
			fmt.Println("Unknown command")
			fmt.Println("Valid commands: upper, lower, cap,title, snake, reverse, exit")
			continue
		}

		if len(part) < 2 && (part[0] == "upper" || part[0] == "UPPER" || part[0] == "lower" || part[0] == "LOWER" ||
			part[0] == "cap" || part[0] == "CAP" || part[0] == "title" || part[0] == "TITLE" || part[0] == "snake" ||
			part[0] == "SNAKE" || part[0] == "reverse" || part[0] == "REVERSE") {
			fmt.Println("Missing text\nUsage: upper <text>")
			continue
		}

		text := part[1:]
		texts := strings.Join(text, " ")

		switch part[0] {
		case "upper":
			fmt.Printf("%s", Upper(texts))
			continue

		case "lower":
			fmt.Printf("%s", Lower(texts))
			continue

		case "cap":
			fmt.Printf("%s", Cap(texts))
			continue

		case "title":
			fmt.Printf("%s", Title(texts))
			continue

		case "snake":
			fmt.Printf("%s", Snake(texts))
			continue

		case "reverse":
			fmt.Printf("%s", Reverse(texts))
			continue

		default:
			fmt.Println("invalid")
			continue
		}

	}
}
