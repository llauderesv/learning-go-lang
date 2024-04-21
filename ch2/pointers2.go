package main

import (
	"fmt"
	"math"
)

// The & ampersand value in the left hand side variable locates the memory address of the pointer
// The * asterisk is used to retrieve the value in the memory address
func main() {
	x := float32(1) // int
	i := &x         // Accessing the memory address of x same instance of the

	fmt.Printf("Value of i: %f\n", *i) // Dereferencing the pointer to retrieve the actual value
	*i = math.Pi
	fmt.Printf("After modifying the pointer i, x would be: %f\n", x)
}

/*
Note: In using a pointer.

You use a pointer if you would like reuse the same value again and again
to point to the previous instance.
*/
