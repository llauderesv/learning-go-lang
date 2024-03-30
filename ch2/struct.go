package main

import "fmt"

// Define a struct named 'Person'
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Create an instance of the 'Person' struct
	person1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Accessing struct fields
	fmt.Println("First Name:", person1.FirstName)
	fmt.Println("Last Name:", person1.LastName)
	fmt.Println("Age:", person1.Age)
}
