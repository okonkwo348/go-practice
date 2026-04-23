package main

import (
	"fmt"
)

// func PrintStr(s string) {

// 	fmt.Print(s)
// }

// func main() {
// 	PrintStr("Hello World!")
// }

// func StrLen(s string) int {
// 	return len(s)
// }

// func main() {
// 	fmt.Println(StrLen("Hello World!"))

// }

// func Swap(a *int, b *int) {
// 	s := *a
// 	t := *b
// 	*b = s
// 	*a = t

// }

// func main() {
// 	a := 0
// 	b := 1
// 	Swap(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// func StrRev(s string) string {
// 	a := []rune(s)
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		a[i], a[j] = a[j], a[i]
// 	}
// 	return string(a)
// }

// func StrRev(s string) string {
// 	a := []rune(s)
// 	c := []rune{}
// 	for i := len(a) - 1; i >= 0; i-- {
// 		c = append(c, a[i])
// 	}
// 	return string(c)
// }

// func main() {
// 	s := "Hello world"
// 	result := StrRev(s)
// 	fmt.Println(result)
// }

// func BasicAtoi(s string) int {
// 	file, _ := strconv.Atoi(s)
// 	return file
// }

func BasicAtoi(s string) int {
	result := 0

	for _, char := range s {
		digit := int(char - '0')
		result = result*10 + digit
	}
	return result
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
