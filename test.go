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

func main() {
	student1 := Student{
		Name:  "Emma",
		Age:   16,
		Grade: 70,
	}

	student2 := Student{"John", 27, 35}

	fmt.Println(student1.introduce())
	fmt.Println(student1.isPassing())

	fmt.Println(student2.introduce())
	fmt.Println(student2.isPassing())
}


// INTERFACE

type Shape interface{
	area() float64
	perimeter() float64
}

type Circle struct{
	radius float64
}

func (c Circle) area() float64{
	return  3.14 * c.radius * c.radius
}

func (c Circle)perimeter()float64{
	return 2 * 3.24 * c.radius 
}

type Rectangle struct{
	width float64
	Height float64
}

func (r Rectangle)area()float64{
	return  r.width * r.Height
}

func (r Rectangle)perimeter()float64{
	return 2*(r.width + r.Height)
}


func printShapeInfo(s Shape){
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

