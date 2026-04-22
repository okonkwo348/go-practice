package main

import "fmt"

// func PrintStr(s string) {

// 	fmt.Print(s)
// }

// func main() {
// 	PrintStr("Hello World!")
// }

func StrLen(s string) int {
	return len(s)
}

func main() {
	fmt.Println(StrLen("Hello World!"))

}
