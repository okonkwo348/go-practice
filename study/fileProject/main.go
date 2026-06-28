package main

import (
	"bufio"
	"fmt"
	"os"
)

func fileExist() bool {
	_, err := os.Stat("contacts.txt")
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func addContact(name string, phone string, email string) {
	if fileExist() {

		file, err := os.OpenFile("contacts.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		file.WriteString("---\n")
		file.WriteString("Name: " + name + "\n")
		file.WriteString("Phone: " + phone + "\n")
		file.WriteString("Email: " + email + "\n")

		fmt.Printf("%s added to contacts!\n", name)

	} else {
		file1, err := os.Create("contacts.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file1.Close()
		file1.WriteString("\n=== Contact Book ===\n")

		file1.WriteString("---\n")
		file1.WriteString("Name: " + name + "\n")
		file1.WriteString("Phone: " + phone + "\n")
		file1.WriteString("Email: " + email + "\n")
	}
}

func viewContacts() {
	_, err := os.Stat("contacts.txt")
	if os.IsNotExist(err) {
		fmt.Println("No contacts found!")
		return
	}

	file, err := os.Open("contacts.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	fmt.Println("=== End of Contacts ===\n")

}

func deleteContacts() {
	_, err := os.Stat("contacts.txt")
	if os.IsNotExist(err) {
		fmt.Println("No contacts found!")
		return
	}
	err = os.Remove("contacts.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("All contacts deleted!")
}

func main() {
	addContact("Okonkwo", "83636376", "okonkwo@gmail.com")
	addContact("emeka", "34578964", "emake@gmail.com")
	addContact("grace", "9877633", "grace@gmail.com")
	//forgot to add this before
	viewContacts()
	deleteContacts()
	viewContacts()
}
