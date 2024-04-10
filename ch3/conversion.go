package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*
		In Go, fmt.Sprintf is a function from the standard fmt package used to format strings.
		It works similarly to fmt.Printf, but instead of printing the formatted string to the standard output (console),
		it returns the formatted string as a string value.
		This allows you to store or manipulate the formatted string as needed.
	*/

	name := "Alice"
	age := 30

	formattedString := fmt.Sprintf("Hello, my name is %s and I'm %d years old", name, age)

	fmt.Println(formattedString)
	fmt.Println(strconv.Itoa(age))
}
