package main

import "fmt"

/*
	An array is a fixed-length sequence of zero or more elements of a particular
	type. Because of their fixed length, arrays are rarely used directly in Go.
	Slices, which can grow and shrink, are much more versatile, but to understand
	slices we must understand arrays first.
*/

func main() {
	// By default, the elements of a new array variable are initially set to the zero value for the element type,
	// which is zero for numbers. We can use array literal to initialize an array with a list of values
	var a [3]int // array of 3 integers

	fmt.Println(a[0]) // print the first element

	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// Print the indices and elements.
	for index, v := range a {
		fmt.Printf("%d %d\n", index, v)
	}

	// Print the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// In an array literal, if an ellipsis "..." appears in place of the length, the array length
	// is determined by the number of initializers.
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	// You can make the array index as value pairs like this:
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	for i, v := range symbol {
		fmt.Printf("%d %s\n", i, v)
	}

}
