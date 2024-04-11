package main

import "fmt"

/*
The function squares returns another function, of type func()
int. A call to squares creates a local variable x, and returns an anonymous function that, each time it is called, increments x and returns its square.
*/

// Function values are not just code but can have state
// The anonymous inner function can access and update the local variables of the enclosing function square.
// JavaScript closure!
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
