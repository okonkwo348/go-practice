package main

import "fmt"

//STRUCT

type Student struct {
	Name  string
	Age   int
	Grade float64
}

func (s Student) introduce() string {
	return fmt.Sprintf("Hi I am %s I am %d years old", s.Name, s.Age)
}

func (s Student) isPassing() bool {
	return s.Grade >= 50
}

// func main() {
// 	student1 := Student{
// 		Name:  "Emma",
// 		Age:   16,
// 		Grade: 70,
// 	}

// 	student2 := Student{"John", 27, 35}

// 	fmt.Println(student1.introduce())
// 	fmt.Println(student1.isPassing())

// 	fmt.Println(student2.introduce())
// 	fmt.Println(student2.isPassing())
// }

// INTERFACE

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

type Rectangle struct {
	width  float64
	Height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.Height
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.Height)
}

func printShapeInfo(s Shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

// func main() {
// 	c := Circle{5}
// 	r := Rectangle{4, 6}

// 	printShapeInfo(c)
// 	printShapeInfo(r)
// }

type Animal interface {
	sound() string
	describe() string
}

type Dog struct {
	Noise  string
	Name   string
	Color  string
	Height string
	nature string
}

func (d Dog) sound() string {
	return fmt.Sprintf("My dog sounds like %s", d.Noise)
}

func (d Dog) describe() string {
	return fmt.Sprintf("My Dog's name is %s, she is %s in complexion, and %s in height, and %s in character", d.Name, d.Color, d.Height, d.nature)
}

type Cat struct {
	Noise  string
	Name   string
	Color  string
	Height string
	nature string
}

func (c Cat) sound() string {
	return fmt.Sprintf("My Cat sounds like %s", c.Noise)
}

func (c Cat) describe() string {
	return fmt.Sprintf("My Cat's name is %s, she is %s in complexion, and %s in height, and %s in character", c.Name, c.Color, c.Height, c.nature)

}

func makeNoise(a Animal) {
	fmt.Println("sound: ", a.sound())
	fmt.Println("description: ", a.describe())
}

func main() {
	d := Dog{Noise: "wo wo", Name: "bingo", Color: "browm and black", Height: "7", nature: "energetic"}
	c := Cat{Noise: "mia", Name: "leis", Color: "white", Height: "4", nature: "gentle"}
	makeNoise(c)
	makeNoise(d)
}
