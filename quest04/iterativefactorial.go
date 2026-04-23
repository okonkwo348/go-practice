package main

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
