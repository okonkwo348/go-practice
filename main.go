package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Check for TWO arguments (input and output)
	// os.Args is a slice of strings.
	// We need 3 elements program (index 0), input (index 2), and output (index 3)
	if len(os.Args) < 3 {
		fmt.Println("Do this: go run . sample.txt result.txt")
		return
	}
	//

	sampleFile := os.Args[1]
	resultFile := os.Args[2]

	//Read the entire file content into a a byte slice
	content, err := os.ReadFile(sampleFile)
	if err != nil {
		// use Printf formatting variables into the string
		fmt.Printf("Can't read file %s: %v\n", sampleFile, err)
		return //Stop if there is an error!
	}

	// 3. Convert bytes to string
	rawText := string(content)

	// 4. THE MAGIC STEP: Apply all your transformations!
	// This calls the big function in your transformations.go file
	modifiedText := ApplyTransformations(rawText)

	// 5. Write the result to the output file
	// 0644 gives the owner read/write permissions
	err = os.WriteFile(resultFile, []byte(modifiedText), 0644)
	if err != nil {
		fmt.Printf("Error writing to output file %s: %v\n", resultFile, err)
		return
	}

	fmt.Printf("Successfully processed %s -> %s\n", sampleFile, resultFile)

	// print the file content by converting the byte slice to a string
	fmt.Println("Content of original file:")
	fmt.Println(string(content))

}
