package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	// "strings"
)

// type BankAccount struct {
// 	Owner   string
// 	Balance float64
// }

// func (b *BankAccount) deposit(amount float64) {
// 	b.Balance += amount
// }

// func (b *BankAccount) withdraw(amount float64) error {
// 	if amount > b.Balance {
// 		return errors.New("insufficient funds")
// 	}
// 	b.Balance -= amount
// 	return nil
// }

// func main() {
// 	account := BankAccount{"emma", 5000}

// 	fmt.Println("initail", account.Balance)

// 	account.deposit(2000)
// 	fmt.Println("After deposit:", account.Balance)

// 	err := account.withdraw(1000)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("After withdrawal:", account.Balance)
// 	}

// 	err2 := account.withdraw(1000)
// 	if err2 != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("After withdrawal:", account.Balance)
// 	}

// 	fmt.Println("after", account.Owner, account.Balance)

// }

// type Student struct {
// 	Name    string
// 	Subject string
// 	Score   float64
// }

// func (s *Student) updateScore(newScore float64) error {
// 	if newScore < 0 || newScore > 100 {
// 		return errors.New("score must be between 0 and 100")
// 	}
// 	s.Score = newScore
// 	return nil
// }

// func (s Student) getGrade() string {
// 	if s.Score >= 90 {
// 		return "A"
// 	}

// 	if s.Score >= 80 {
// 		return "B"
// 	}

// 	if s.Score >= 70 {
// 		return "C"
// 	}

// 	if s.Score >= 50 {
// 		return "D"
// 	}
// 	return "F"
// }

// func main() {
// 	student1 := Student{"peter", "maths", 50.3}
// 	student2 := Student{"emma", "maths", 78}

// 	fmt.Println("initial", student1.Score)

// 	err := student1.updateScore(23)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	fmt.Println("After upgrade: ", student1.Score)

// 	err2 := student2.updateScore(150)
// 	if err2 != nil {
// 		fmt.Println("Error:", err2)
// 	}
// 	fmt.Println("After: ", student2.Score)

// 	fmt.Println(student1.getGrade())
// 	fmt.Println(student2.getGrade())

// }

// Scanf

// func main() {
// 	var name string
// 	var age int
// 	fmt.Print("What is your name:")
// 	fmt.Scanf("%s", &name)
// 	fmt.Print("What is your age:")
// 	fmt.Scanf("%d", &age)

// 	fmt.Printf("Hello %s, you are %d years old", name, age)
// }

// bufio

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter your city: ")
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city)

	fmt.Printf("%s lives in %s\n", name, city)
}

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	fmt.Print("Enter your full name: ")
// 	scanner.Scan()
// 	name := scanner.Text()

// 	fmt.Print("Enter your city: ")
// 	scanner.Scan()
// 	city := scanner.Text()
// 	fmt.Printf("%s lives in %s\n", name, city)
// }
