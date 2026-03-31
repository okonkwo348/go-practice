package main

// creating a file using os.Create

// func main() {
// 	file, err := os.Create("hello.txt")
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 		return
// 	}
// 	defer file.Close()

// 	fmt.Println("Name of file: ", file.Name())

// }

// creating a file using os.OpenFile
// func main() {
// 	file, err := os.OpenFile("text.txt", os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Println("error :", err)
// 		return
// 	}

// 	defer file.Close()

// 	fmt.Println("Name of file: ", file.Name())

// }

// Writing files

// 1. method One: file.WriteString()
// func main() {
// 	file, err := os.Create("hell.txt")
// 	if err != nil {
// 		fmt.Println("error:", err)
// 		return
// 	}

// 	defer file.Close()

// 	file.WriteString("Name: Okonkwo Emmanuel\n")
// 	file.WriteString("Age: 23")

// 	fmt.Printf("%v file created sucessfully", file.Name())
// }

//Method 2: Fprintf()
// func main() {
// 	file, err := os.Create("text.txt")
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 		return
// 	}

// 	defer file.Close()

// 	name := "Okonkwo Emma"
// 	age := 23

// 	fmt.Fprintf(file, "Name: %s\nage: %v\n", name, age)
// 	// fmt.Fprintf(file, "age: %v\n", age)
// }

// Method 3: os.WriteFile
// func main() {
// 	content := "Name: Okonkwo\nAge: 23"
// 	err := os.WriteFile("hello.txt", []byte(content), 0644)
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 		return
// 	}
// 	// close file is handled automatically

// 	fmt.Println("File written successfully")

// }

// Exercise
// func main() {
// 	file, err := os.Create("text.txt")
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 		return
// 	}
// 	defer file.Close()

// 	Name := "okonkwo"
// 	Age := 34
// 	Course := "python"
// 	Grade := 89

// 	fmt.Fprintf(file, "Name: %s\nAge: %d\nCourse: %s\nGrade: %d", Name, Age, Course, Grade)
// }
