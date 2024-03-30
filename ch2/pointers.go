package main

import (
	"fmt"
)

func main() {
	x := 1

	var i *int //
	i = &x     // passing the address of variable x
	*i = 2     // Dereferencing the variable

	fmt.Printf("The value of i: %d is equivalent to x: %d\n", *i, x)

	fmt.Println(f() == f()) // false
	// Each call of f returns a distinct value:

	v := 1

	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)

	medals := []string{"gold", "silver", "bronze"}
	medals[0] = "gold"
}

var p = f()

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p

	return *p
}
