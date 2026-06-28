package main

import (
	"fmt"
	"os"
)

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

// Method 4: OpenFile  with O_CREATE and O_WRONLY
// func main() {
// 	// overwrites existing content
// 	file, err := os.OpenFile("student.txt", os.O_WRONLY, 0644)

// 	// adds to existing content
// 	// file, err := os.OpenFile("student.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 		return
// 	}

// 	defer file.Close()

// 	file.WriteString("Hi")

// 	fmt.Println("File written successful")

// }

// READING A FILE
// Method 1: os.ReadFile()
// func main() {
// 	content, err := os.ReadFile("student.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println(content)
// }

// os.Open + bufio.Scanner
// read the file line by line - better for large files
// func main() {
// 	file, err := os.Open(("student.txt"))
// 	if err != err {
// 		fmt.Println("error: ", err)
// 		return
// 	}

// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// }

//Method 2: os.Open + bufio.NewReader

// func main() {
// 	file, err := os.Open(("student.txt"))
// 	if err != err {
// 		fmt.Println("error: ", err)
// 		return
// 	}

// 	defer file.Close()

// 	reader := bufio.NewReader(file)

// 	for {
// 		line, err := reader.ReadString(':')
// 		line = strings.TrimSpace(line)

// 		if line != "" {
// 			fmt.Println(line)
// 		}

// 		if err == io.EOF { // EOF means End Of File — no more lines
// 			break
// 		}
// 	}
// }

// func createDiary(filename string) {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 		return
// 	}
// 	file.WriteString("=== My Diary ===\n")
// 	fmt.Println("Diary created!")
// }

// func addEntry(filename string, entry string) {
// 	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 		return
// 	}
// 	defer file.Close()

// 	file.WriteString(entry + "\n")

// 	fmt.Printf("%s added to diary!\n", entry)
// }

// func readDiary(filename string) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 		return
// 	}
// 	defer file.Close()
// 	fmt.Println()

// 	scanner := bufio.NewScanner(file)

// 	newLine := 1
// 	for scanner.Scan() {
// 		fmt.Printf("Line %d: %s\n", newLine, scanner.Text())
// 		newLine++
// 	}

// 	fmt.Println("=== End of Diary ===")
// }

// func main() {
// 	createDiary("diary.txt")
// 	addEntry("diary.txt", "Today I started learning Go! added to diary!")
// 	addEntry("diary.txt", "I learned file handling today! added to diary!")
// 	addEntry("diary.txt", "Go is becoming my favorite language! added to diary!")
// 	readDiary("diary.txt")
// }

//Checking if a File Exist

// func fileExists(filename string) bool {
// 	_, err := os.Stat(filename)
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }

// func checkFile(filename string) {
// 	if fileExists(filename) {
// 		fmt.Printf("%s exists!\n", filename)
// 		file, err := os.ReadFile(filename)
// 		if err != nil {
// 			fmt.Println("Error: ", err)
// 			return
// 		}
// 		fmt.Println(string(file))

// 	} else {
// 		fmt.Printf("%s not found — creating it\n", filename)
// 		content, err := os.Create(filename)
// 		if err != nil {
// 			fmt.Println("Error :", err)
// 			return
// 		}
// 		defer content.Close()
// 		content.WriteString("New file created!\n")
// 	}
// }

// func main() {
// 	checkFile("diary.txt")
// 	checkFile("now.txt")

// }

//Exercise

// func fileExists(filename string) bool {
// 	_, err := os.Stat(filename)
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }
// func createRecord(filename string, name string, course string) {
// 	if fileExists(filename) {
// 		fmt.Println("Record already exists!")
// 		return
// 	} else {
// 		file, err := os.Create(filename)
// 		if err != nil {
// 			fmt.Println("Error: ", err)
// 			return
// 		}
// 		defer file.Close()

// 		file.WriteString("=== Student Record ===\n")
// 		file.WriteString("Name: " + name + "\n")
// 		file.WriteString("Course: " + course + "\n")

// 		fmt.Printf("Record created for %s!", name)

// 	}
// }

// func addScore(filename string, subject string, score float64) {
// 	if fileExists(filename) {
// 		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 		if err != nil {
// 			fmt.Println("Error: ", err)
// 			return
// 		}
// 		defer file.Close()

// 		fmt.Fprintf(file, "%s: %.2f\n", subject, score)
// 		fmt.Printf("%s score added!\n", subject)
// 	} else {
// 		fmt.Println("Record not found!")
// 		return
// 	}
// }

// func viewRecord(filename string) {
// 	if fileExists(filename) {
// 		file, err := os.Open(filename)
// 		if err != nil {
// 			fmt.Println("error: ", err)
// 			return
// 		}
// 		defer file.Close()

// 		scanner := bufio.NewScanner(file)

// 		newLine := 1
// 		for scanner.Scan() {
// 			fmt.Printf("Line %d: %s\n", newLine, scanner.Text())
// 			newLine++
// 		}
// 		fmt.Println("=== End of Record ===")
// 	} else {
// 		fmt.Println("Record not found!")
// 		return
// 	}
// }

// func main() {
// 	createRecord("student.txt", "Okonkwo", "Physics")
// 	createRecord("student.txt", "Okonkwo", "Physics")
// 	addScore("student.txt", "maths", 50)
// 	addScore("student.txt", "english", 90)
// 	addScore("student.txt", "science", 30)
// 	viewRecord("student.txt")
// }

//Deleting FILE

// func main() {
// 	err := os.Remove("hell.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("File deleted successfully")
// }

// Always check if the file exists beforrrre deleting:
// func deleteFile(filename string) {
// 	_, err := os.Stat(filename)
// 	if os.IsNotExist(err) {
// 		fmt.Printf("%s nnnnot found - nothing to delete\n", filename)
// 		return
// 	}
// 	err = os.Remove(filename)
// 	if err != nil {
// 		fmt.Println("Error deleting file:", err)
// 		return
// 	}
// 	fmt.Printf("%s deleted sucessfully!\n", filename)
// }

// func main() {
// 	deleteFile("hello.txt")
// 	deleteFile("hell.txt")
// }

//Exercise

func safeDelete(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Printf("%s not found!\n", filename)
		return
	}
	err = os.Remove(filename)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Printf("%s deleted!", filename)
}

func main() {
	file, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	file.WriteString("This is a lesson of file handling")

	safeDelete("temp.txt")
	safeDelete("temp.txt")
}
