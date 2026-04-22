package main

import "fmt"

func PointOne(n ***int) {
	***n = 1
}

func main() {
	a := 0
	b := &a
	c := &b

	PointOne(&c)
	fmt.Println(a)
}
