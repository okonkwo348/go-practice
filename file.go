package main

import (
	"fmt"
	"os"
)

// creating a file using os.Create

// func main() {
// 	file, err := os.Create("hello.txt")
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 		return
// 	}
// 	defer file.Close()

// 	fmt.Println("Name of file: ", file.Name())

// }

// creating a file using os.OpenFile
// func main() {
// 	file, err := os.OpenFile("text.txt", os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Println("error :", err)
// 		return
// 	}

// 	defer file.Close()

// 	fmt.Println("Name of file: ", file.Name())

// }

// Writing files

// 1. method One: file.WriteString()
func main() {
	file, err := os.Create("hell.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer file.Close()

	file.WriteString("Name: Okonkwo Emmanuel\n")
	file.WriteString("Age: 23")

	fmt.Printf("%v file created sucessfully", file.Name())
}
