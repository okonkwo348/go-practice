package main

import (
	"fmt"
	"strings"
)

// Your exercise:
// Write a program that:

// Declares a full name "Okonkwo Peter Emmanuel"
// Prints the total length
// Prints the first 7 characters
// Prints the last 9 characters
// Prints the full name using fmt.Sprintf with first,
// middle and last name as separate variables

// func main() {
// 	name := "Okonkwo Peter Emmanuel"

// 	fmt.Println(len(name))
// 	fmt.Println(name[:7])
// 	fmt.Println(name[len(name)-9:])

// 	first := "Okonkwo"
// 	middle := "Peter"
// 	last := "Emmanuel"
// 	fullName := fmt.Sprintf("%s %s %s", first, middle, last)
// 	fmt.Println(fullName)
// }

// Exercise — Name Formatter
// Write a program that takes this messy full name string:
// goname := "okonkwo peter emmanuel"
// And using only basic string operations:

// Print the total number of characters including spaces
// Print the first 3 characters of each name using slicing
// Build a new string called initials that contains "O.P.E"
// using only individual characters from the string
// Print the middle name by slicing out "peter" from the full string
// Print the full name length without spaces — hint: you'll need to count manually using a loop

// func main() {
// 	name := "okonkwo peter emmanuel"

// 	fmt.Println(len(name))

// 	s := strings.Fields(name)
// 	s[0] = "Oko"
// 	s[1] = "Pet"
// 	s[2] = "Em"

// 	fmt.Println(strings.Join(s, " "))

// 	slice := strings.Fields(name)
// 	fmt.Println(slice[1])

// 	initials := strings.Fields(name)
// 	initials[0] = "O."
// 	initials[1] = "P."
// 	initials[2] = "E"

// 	count := 0
// 	for _, char := range name {
// 		if string(char) == " " {
// 			continue
// 		}
// 		count++
// 	}
// 	fmt.Println(count)

// }

//CORRECTION

// func main() {
// 	name := "okonkwo peter emmanuel"

// 	fmt.Println(len(name))

// 	fmt.Println(name[0:3], name[8:11], name[14:17])
// 	intials := fmt.Sprintf("%s.%s.%s", string(name[0]), string(name[8]), string(name[14]))
// 	fmt.Println(intials)
// 	fmt.Println(name[8:13])

// 	count := 0
// 	for _, char := range name {
// 		if string(char) == " " {
// 			continue
// 		}
// 		count++
// 	}
// 	fmt.Println(count)
// }

// Your exercise:
// Write a program that takes this string:
// gosentence := "  i love java. java is good but go is better than java.  "
// And using the strings package:

// Trim the spaces and print it
// Print it in uppercase
// Count how many times "java" appears
// Replace all "java" with "Go"
// Split the sentence into words and print the total word count
// Check if the sentence contains "go" and print the result

//Use stings Package
// func main() {
// 	sentence := "  i love java. java is good but go is better than java.  "

// 	fmt.Println(strings.TrimSpace(sentence))
// 	fmt.Println(strings.ToUpper(sentence))
// 	fmt.Println(strings.Count(sentence, "java"))
// 	fmt.Println(strings.ReplaceAll(sentence, "java", "Go"))
// 	split := strings.Fields(sentence)
// 	fmt.Println(len(split))
// 	fmt.Println(strings.Contains(sentence, "go"))
// }

// // strconv Package
// func main() {
// 	ageStr := "24"
// 	priceStr := "49.99"
// 	isActiveStr := "true"

// 	num, err := strconv.Atoi(ageStr)
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 		return
// 	}

// 	flo, err1 := strconv.ParseFloat(priceStr, 64)
// 	if err1 != nil {
// 		fmt.Println("Error: ", err1)
// 		return
// 	}

// 	bol, err2 := strconv.ParseBool(isActiveStr)
// 	if err2 != nil {
// 		fmt.Println("Error: ", err2)
// 		return
// 	}

// 	fmt.Println(num + 6)
// 	fmt.Println(flo * 3)
// 	fmt.Println(bol)

// 	fmt.Printf("%s | %s | %s", strconv.Itoa(num), strconv.FormatFloat(flo, 'f', 2, 64), strconv.FormatBool(bol))

// }package main

// func main() {

// 	username := "  Okonkwo  "
// 	ageStr := "17"
// 	balanceStr := "250.50"
// 	isActiveStr := "false"

// 	username = strings.TrimSpace(strings.ToUpper(username))

// 	age, err := strconv.Atoi(ageStr)
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		return
// 	}

// 	if age < 18 {
// 		fmt.Println("User must be 18 or older")
// 		return
// 	}

// 	balance, err2 := strconv.ParseFloat(balanceStr, 64)
// 	if err2 != nil {
// 		fmt.Println("error: ", err2)
// 		return
// 	}

// 	balance = balance + 500
// 	fmt.Println(strconv.FormatFloat(balance, 'f', 2, 64))

// 	active, err3 := strconv.ParseBool(isActiveStr)
// 	if err3 != nil {
// 		fmt.Println("Error: ", err3)
// 		return
// 	}
// 	isNowActive := !active
// 	if isNowActive == true {
// 		fmt.Println("Account activated!")
// 	}

// 	summary := []string{"=== User Profile ===",
// 		"Username: " + username,
// 		"Age: " + strconv.Itoa(age),
// 		"Balance: " + strconv.FormatFloat(balance, 'f', 2, 64),
// 		"Active: " + strconv.FormatBool(isNowActive)}
// 	fmt.Println(strings.Join(summary, "\n"))

// }

// strings.Bulider

func buildReport(title string, items []string) string {
	var now strings.Builder

	now.WriteString(fmt.Sprintf("=== %s===\n", title))

	for i, char := range items {
		now.WriteString(fmt.Sprintf("%d. %s\n", i+1, char))
	}

	now.WriteString("--- End of Report ---")

	return now.String()

}

func main() {
	title := "My Shopping List"
	items := []string{"Apples", "Milk", "Bread", "Eggs"}
	fmt.Println(buildReport(title, items))

}
