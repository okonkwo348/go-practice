package main

import "fmt"

func main() {
	name := "Okonkwo Peter Emmanuel"

	fmt.Println(len(name))
	fmt.Println(name[:7])
	fmt.Println(name[len(name)-9:])

	first := "Okonkwo"
	middle := "Peter"
	last := "Emmanuel"
	fullName := fmt.Sprintf("%s %s %s", first, middle, last)
	fmt.Println(fullName)
}
