package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	
start:
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Println(">")
	if !scanner.Scan(){
		goto start
	}

	line:=scanner.Text()
	part:=strings.Fields(line)

	if part[0]=="quit"{
		fmt.Println("exit")
	}

	if part[0]=="HELP"{
		fmt.Println("To add type: <sum> <num1> <num2>")
		fmt.Println("To substrate type: <sub> <num1> <num2>")
		fmt.Println("To mulply type: <mul> <num1> <num2>")
		fmt.Println("To divide type: <div> <num1> <num2>")
	}

	if len(part)>0 && part[0]=="add" || part[0]=="sub" || part[0]=="mul" || part[0]=="div"{
		switch part[0] {
		case "add":
			fmt.Println(add(part[1],part[2]))
			goto start
		case "sub":
			fmt.Println(sub(part[1],part[2]))
			goto start
		case "mul":
			fmt.Println(mul(part[1],part[2]))
			goto start
		case "div":
			if part[2]=="0"{
				fmt.Println("can't divide zero")
				goto start
			}
			fmt.Println(div(part[1],part[2]))
			goto start
		default:
		fmt.Println("usage: <sum><sub><mul><div> <num1> <num2>")
		goto start
	    }
	}
}