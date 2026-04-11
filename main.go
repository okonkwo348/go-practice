package main
import (
	"fmt"
	"os"
)
func validateInput(s string)([]string, error){
	var result []string
	for i := 0, i < len(s); i++{
		if !(s[i] <= 32 && s[i] >= 126 || s[i] == ' '){
			return nil, error.new("invalid input")
		}
		result = append(result, s[i])
	}
	return result, nil
}

func main(){
	input := os.Args

	part := string.Split(validateInput(input),"\n")

	// Chang if argument is provided
	if len(input) != 2{
		return 
	}

	if input[1] == ""{
		return fmt.Println("\n")
	}


}