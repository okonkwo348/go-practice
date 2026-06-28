package main

import (
	"fmt"
	// "os"
	"strings"
)

func lastChar(s string) string {
	if len(s) == 0 { return "" }
	return s[len(s)-1:]
}

func capitalize(s string) string {
	return strings.ToUpper(s)
}

func main(){
    word:=""
    var index int
    fmt.Println("Enter a word: ")
    fmt.Scan(&word)

start:
    fmt.Println("Select one of the operations below: ")
    fmt.Println("lastChar")
    fmt.Println("capitalize")
    fmt.Println("deleteIndex")

    var operations string
    fmt.Scan(&operations)
    switch operations{
    case "lastChar":
        fmt.Println(lastChar(word))
    case "capitalize":
        fmt.Println(capitalize(word))
    case "deleteIndex":
        fmt.Println("write an index")
        fmt.Scan(&index)
        if len(word)==index || index > len(word) || len(word)==0{
            fmt.Println("error: index greater than length of the word")
            goto start
            
        }
        fmt.Println(word[:index] + word[index+1:])
    default:
        fmt.Println("invalid Oprations")
        goto start
    }
}