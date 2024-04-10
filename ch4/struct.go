package main

import (
	"fmt"
	"time"
)

/*
A struct is an aggregate data type that groups together zero or
more named values of arbitrary types as a single entity.

Each value is called field.

Struct field can be accessed via dot notation
*/

type Employee struct {
	ID            int
	Name, Address string // Consecutive fields of the same type may be combined
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

type Point struct{ X, Y int }

func main() {
	var dilbert Employee

	dilbert.Salary -= 5000

	position := &dilbert.Position // take its address and access it through a pointer

	*position = "Senior "
	fmt.Printf("%s\n", *position)
	fmt.Printf("Dilbert position: %s\n", *&dilbert.Position)
	dilbert.Position = "Senior Software Engineer"
	fmt.Printf("After position: %s\n", *position)

	p := Point{1, 2}

	fmt.Printf("Value of X: %d, and the value of Y: %d\n", p.X, p.Y)
	fmt.Printf("Value of X: %d, and the value of Y: %d\n", p.X, p.Y)

	scaleP := Scale(&p, 5)
	fmt.Printf("Scale X: %d, Scale Y: %d", scaleP.X, scaleP.Y)

	fmt.Printf("Value of X: %d, and the value of Y: %d\n", p.X, p.Y)
	fmt.Printf("Value of X: %d, and the value of Y: %d\n", p.X, p.Y)
}

// The function EmployeeByID returns a pointer to an Employee struct.
// We can use the dot notation its fields:
func EmployeeByID(id int) *Employee {
	var employee Employee

	return &employee
}

func Scale(p *Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}
