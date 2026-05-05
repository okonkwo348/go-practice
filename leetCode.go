package main

import (
	"fmt"
)



func gcdOfStrings(str1 string, str2 string) string {
	// s := ""
	// for i := 0; i < len(str1); i++ {
	// 	for j := 0; j < len(str2); j++ {
	// 		if i == j {
	// 			if !strings.ContainsAny(s, string(str2[i])) {
	// 				s += string(str1[i])
	// 			}
	// 		}

	// 	}
	// }

	// if len(str1)%len(s) == 0 && len(str2)%len(s) == 0 {
	// 	return s
	// }
	// return ""
	if str1+str2 != str2+str1 {
		return ""
	}

	length := GCD(len(str1), len(str2))
	return str1[:length]

}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(gcdOfStrings("ABAB", "AB"))
	fmt.Println(gcdOfStrings("NANANAF", "NANA"))

	fmt.Println(gcdOfStrings("AAAA", "AA"))
}
