package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func main(){
// 	name:="Okonkwo Peter Emmanuel"

// 	fmt.Println(len(name))
// 	fmt.Println(name[:8])
// 	fmt.Println([:10])

// 	first:="Okonkwo"
// 	middle:="Peter"
// 	last:="Emmanuel"
// 	fmt.Sprintf("%s %s %s",first,middle,last)

// }

func main() {

	username := "  Okonkwo  "
	ageStr := "17"
	balanceStr := "250.50"
	isActiveStr := "false"

	username = strings.TrimSpace(strings.ToUpper(username))

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if age < 18 {
		fmt.Println("User must be 18 or older")
		return
	}

	balance, err2 := strconv.ParseFloat(balanceStr, 64)
	if err2 != nil {
		fmt.Println("error: ", err2)
		return
	}

	balance = balance + 500
	fmt.Println(strconv.FormatFloat(balance, 'f', 2, 64))

	active, err3 := strconv.ParseBool(isActiveStr)
	if err3 != nil {
		fmt.Println("Error: ", err3)
		return
	}
	isNowActive := !active
	if isNowActive == true {
		fmt.Println("Account activated!")
	}

	summary := []string{"=== User Profile ===",
		"Username: " + username,
		"Age: " + strconv.Itoa(age),
		"Balance: " + strconv.FormatFloat(balance, 'f', 2, 64),
		"Active: " + strconv.FormatBool(isNowActive)}
	fmt.Println(strings.Join(summary, "\n"))

}
