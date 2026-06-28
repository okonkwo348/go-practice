// package main

// import "fmt"

// func main() {
// 	name := "Okonkwo"
// 	year := 2000
// 	fmt.Printf("name: %s\nyear: %d\n", name, year)
// }

// package main

// import "fmt"

// func main() {
// 	// Declare these three ways:
// 	// 1. Short declaration (:=)
// 	a := 1
// 	b := 1.1
// 	c := true
// 	// 2. Explicit type (var x int = 5)
// 	var x int = 5
// 	var y float64 = 5.7
// 	var z bool = false
// 	// 3. Zero value (var x int — no assignment)
// 	var l int
// 	var m float64
// 	var n bool

// 	// Then print all three and observe Go's zero values
// 	fmt.Println(a, b, c)
// 	fmt.Println(x, y, z)
// 	fmt.Println(l, m, n)
// }

// Exercise 1 — Warm up
// Write a function isAdult that takes an age int and returns a bool. In main, use it to print whether a person is an adult or not.
// go// Expected output for age 20:
// "You are an adult: true"

// Expected output for age 15:
// "You are an adult: false"

// package main

// import "fmt"

// func isAdult(age int) bool {
// 	return age >= 18
// }
// func main() {

// 	fmt.Printf("you are an adult: %v\n", isAdult(20))
// 	fmt.Printf("you are an adult: %v", isAdult(15))
// }

// Exercise 2 — Combining bools
// Write a function canAccessSystem that takes two bools — isLoggedIn and hasPermission — and returns true only if both are true. Test all four combinations in main.
// go// Expected output:
// true, true   → "Access granted"
// true, false  → "Access denied"
// false, true  → "Access denied"
// false, false → "Access denied"

// func canAccessSystem(isLoggedIn bool, hasPermission bool) string {
// 	if isLoggedIn && hasPermission {
// 		return "Access granted"
// 	}
// 	return "Access denied"
// }
// func main() {

// 	fmt.Printf("%v\n", canAccessSystem(true, true))
// 	fmt.Printf("%v\n", canAccessSystem(true, false))
// 	fmt.Printf("%v\n", canAccessSystem(false, true))
// 	fmt.Printf("%v\n", canAccessSystem(false, false))
// }

// Exercise 3 — Bool + logic together
// Write a function classifyNumber that takes an int and prints:

// "positive and even" if the number is both
// "positive and odd" if positive but odd
// "negative" if the number is negative
// "zero" if it's 0

// go// Hints:
// // - % is the modulo operator (n % 2 == 0 means even)
// // - You'll need && (AND) and possibly || (OR)

// func classifyNumber(n int) string {
// 	if n%2 == 0 && n > 0 {
// 		return "positive and even"
// 	} else if n%2 != 0 && n > 0 {
// 		return "positive and odd"
// 	} else if n < 0 {
// 		return "negative"
// 	} else if n == 0 {
// 		return "zero"
// 	}
// 	return "invalid"
// }

// func main() {
// 	fmt.Println(classifyNumber(3))
// 	fmt.Println(classifyNumber(6))
// 	fmt.Println(classifyNumber(-3))
// 	fmt.Println(classifyNumber(0))
// }

//	ERROR HANDLING

// Exercise — divide function
// gopackage main

// import (
//     "errors"
//     "fmt"
// )

// func divide(a float64, b float64) (float64, error) {
//     // your logic here
// }

// func main() {
//     result, err := divide(10, 2)
//     // handle err, then print result

//     result2, err2 := divide(10, 0)
//     // handle err2, then print result2
// }
// Your task:

// If b is 0 return an error using errors.New("cannot divide by zero")
// If b is valid, return the result of a / b and nil for the error
// In main handle both results — print the result if no error, print the error if there is one

// Checklist before you submit:

//  Does your function return (float64, error)?
//  Did you guard against b == 0?
//  Did you check err != nil in main before printing results?

// Take your time and paste your code when ready! 🚀Exercise — divide function
// gopackage main

// import (
//     "errors"
//     "fmt"
// )

// func divide(a float64, b float64) (float64, error) {
//     // your logic here
// }

// func main() {
//     result, err := divide(10, 2)
//     // handle err, then print result

//     result2, err2 := divide(10, 0)
//     // handle err2, then print result2
// }
// Your task:

// If b is 0 return an error using errors.New("cannot divide by zero")
// If b is valid, return the result of a / b and nil for the error
// In main handle both results — print the result if no error, print the error if there is one

// Checklist before you submit:

//  Does your function return (float64, error)?
//  Did you guard against b == 0?
//  Did you check err != nil in main before printing results?

// Take your time and paste your code when ready! 🚀
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func divide(a float64, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, errors.New("b can not be 0")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	result, err := divide(10, 2)
// 	if err != nil {
// 		fmt.Println("error", err)
// 		return
// 	}
// 	fmt.Println("result", result)

// 	result2, err2 := divide(10, 0)
// 	if err2 != nil {
// 		fmt.Println("error", err2)
// 		return
// 	}
// 	fmt.Println("result", result2)
// }

// Try writing your own — a function called login that takes a username string and password string and returns (string, error). Rules:

// if username is empty → error
// if password is less than 6 characters → error
// otherwise → return "welcome " + username, nil

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func login(username string, password string) (string, error) {
// 	if username == "" {
// 		return "", errors.New("username is empty")
// 	}
// 	if len(password) < 6 {
// 		return "", errors.New("password is less than 6 char")
// 	}
// 	return "welcome " + username, nil
// }

// func main() {
// 	result, err := login("emmanuel", "emmapeter")
// 	if err != nil {
// 		fmt.Println("error", err)
// 		return
// 	}
// 	fmt.Println("result", result)

// 	result, err2 := login("emmanuel", "emma")
// 	if err2 != nil {
// 		fmt.Println("error", err2)
// 		return
// 	}
// 	fmt.Println("result", result)

// 	result, err3 := login("", "emmapeter")
// 	if err3 != nil {
// 		fmt.Println("error", err3)
// 		return
// 	}
// 	fmt.Println("result", result)
// }

// STRUCT

// Your exercise 1:
// Create a Student struct with these fields:

// Name string
// Age int
// Grade float64

// Then add two methods:

// introduce() — returns "Hi I am {name} and I am {age} years old"
// isPassing() — returns true if Grade is 50 and above, false if below

// In main create two students — one passing, one failing — and call both methods on each.

// package main

// import (
// 	"fmt"
// )

// type Student struct {
// 	Name  string
// 	Age   int
// 	Grade float64
// }

// func (s Student) introduce() string {
// 	return fmt.Sprintf("Hi I am %s and I am %d years old", s.Name, s.Age)
// }

// func (s Student) isPassing() bool {
// 	return s.Grade >= 50

// }

// Your next exercise:
// Add a method updateGrade to your Student struct that takes a newGrade float64 and updates the student's grade.
// Then  in main print the grade before and after calling it.

// func (s *Student) updateGrade(newGrade float64) {
// 	s.Grade = newGrade
// }

// func main() {
// 	score := Student{"okonkwo", 20, 70.0}
// 	fmt.Println(score.introduce())
// 	fmt.Println(score.isPassing())

// 	score2 := Student{"joy", 43, 20.0}
// 	fmt.Println(score2.introduce())
// 	fmt.Println(score2.isPassing())

// 	fmt.Println(score2.Grade)

// 	score2.updateGrade(50.5)
// 	fmt.Println(score2.Grade)
// }

// // INTERFACE
// Your exercise:
// Create an interface called Animal with two methods:

// sound() string — returns the sound the animal makes
// describe() string — returns a short description

// Then create two structs — Dog and Cat — that satisfy the interface.
// Write a function makeNoise(a Animal) that calls both methods and prints them. Test with both animals in main.

package main

type Animal interface {
	sound() string
	describe() string
}

type Dog struct {
	sound    string
	describe string
}

type Cat struct {
	sound    string
	describe string
}

func (a Animal) makeNoise()
