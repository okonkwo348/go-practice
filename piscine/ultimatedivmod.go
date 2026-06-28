package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	*a = *a / *b
	*b = *a % *b
}

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
