package main

import "fmt"

/* Write an iterative function that returns the factorial of the int passed as parameter.

Errors (non possible values or overflows) will return 0.*/
// func IterativeFactorial(nb int) int {
// 	num := 1
// 	for i := 1; i <= nb; i++ {
// 		num *= i
// 	}
// 	return num
// }

// func main() {
// 	arg := 4
// 	fmt.Println(IterativeFactorial(arg))
// }

/* Write a recursive function that returns the factorial of the int passed as parameter.

Errors (non possible values or overflows) will return 0.*/

// func RecursiveFactorial(nb int) int {

// 	num := 1
// 	i := 0

// 	for i < nb {
// 		i++
// 		num *= i
// 	}
// 	return num
// }

// func main() {
// 	arg := 4
// 	fmt.Println(RecursiveFactorial(arg))
// }

/* Write an iterative function that returns the value of nb to the power of power.

Negative powers will return 0. Overflows do not have to be dealt with.*/

// func IterativePower(nb int, power int) int {
// 	num := 1
// 	for i := 0; i < power; i++ {
// 		num *= nb
// 	}
// 	return num
// }

// func main() {
// 	fmt.Println(IterativePower(3, 4))
// }

/* Write a recursive function that returns the value of nb to the power of power.

Negative powers will return 0. Overflows do not have to be dealt with.

for is forbidden for this exercise.*/

// func RecursivePower(nb int, power int) int {
// 	num := 1

// 	i := 0
// 	for i < power {
// 		i++
// 		num *= nb
// 	}
// 	return num
// }

// func main() {
// 	fmt.Println(RecursivePower(4, 3))
// }

/* Write a recursive function that returns the value at the position index in the fibonacci sequence.

The first value is at index 0.

The sequence starts this way: 0, 1, 1, 2, 3 etc...

A negative index will return -1.

for is forbidden for this exercise.*/

// func Fibonacci(index int) int {
// 	num := []int{4, 1, 7, 2, 3, 8}
// 	x := 0
// 	for i := 0; i < len(num); i++ {
// 		if i == index {
// 			x += num[i]
// 			break
// 		}
// 	}
// 	return x
// }

// func main() {
// 	arg1 := 5
// 	fmt.Println(Fibonacci(arg1))
// }

func IsPrime(nb int) bool {

	if nb%2 == 0 || nb%3 == 0 || nb%5 == 0 || nb%7 == 0 || nb%11 == 0 {
		return false
	}
	return true

}

func main() {
	fmt.Println(IsPrime(110))
	fmt.Println(IsPrime(9))
}
