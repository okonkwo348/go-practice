package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	var contain []string

// 	file, err := os.Create("text.txt")
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		return
// 	}
// 	defer file.Close()

// 	file.WriteString("My name is Emmanuel, i came from lagos.\n")
// 	file.WriteString("I feel mlike go is the best language.\n")
// 	file.WriteString("i am learning slice at the moment.")

// 	file1, err1 := os.Open("text.txt")
// 	if err1 != nil {
// 		fmt.Println("Error: ", err1)
// 		return
// 	}
// 	defer file1.Close()

// 	scanner := bufio.NewScanner(file1)
// 	for scanner.Scan() {
// 		part := scanner.Text()
// 		contain = append(contain, part)
// 	}

// 	fmt.Println(len(contain))
// 	fmt.Println(contain[1])

// }
func loopEach(s []string) string {

	result := []string{}
	for _, v := range s {
		result = append(result, v)
	}
	return strings.Join(result, "\n")
}
func main() {
	var big [][]string
	charA := []string{" _ ", "| |", "|_|"}
	charB := []string{" _ ", "| }", "|_}"}
	charC := []string{" _ ", "| ", "|_ "}

	big = append(big, charA, charB, charC)

	result := []string{}
	for _, v := range s {
		result = append(result, v)
	}


	fmt.Println(loopEach(big[]))

}
