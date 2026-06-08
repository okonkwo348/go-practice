package main

import (
	"fmt"
	"strings"
)

func Split(s string) []string {
	return strings.Fields(s)
}

func Split2(s string) []string {
	word := strings.Fields(s)
	result := []string{}
	for i := 0; i < len(word); i++ {

		if strings.ContainsAny(word[i], ",") {
			c := ""
			for _, r := range word[i] {
				if string(r) == "," {
					c = " " + string(r)
				}
				c += string(r)
			}
			d := strings.Split(c, " ")
			for i := 0; i < len(d); i++ {
				result = append(result, d[i])
			}
		}

		if strings.ContainsAny(word[i], "!") {
			k := ""
			for _, r := range word[i] {
				if string(r) == "!" {
					k = " " + string(r)
				}
				k += string(r)
			}
			b := strings.Split(k, "!")
			for i := 0; i < len(b); i++ {
				result = append(result, b[i])
			}
		}

		if strings.ContainsAny(word[i], "?") {
			v := ""
			for _, r := range word[i] {
				if string(r) == "?" {
					v = " " + string(r)
				}
				v += string(r)
			}
			a := strings.Split(v, "?")
			for i := 0; i < len(a); i++ {
				result = append(result, a[i])
			}
		}

		result = append(result, word[i])

	}

	return result

}

func main() {
	fmt.Println(Split("Hello, world! How are you?"))
	fmt.Println(Split2("Hello, world! How are you?"))
}
