package main

import (
	"fmt"
	"math"
)

func main() {
	// When a sequence of constant is declared as a group, the right
	// hand side expression maybe omitted for all but the first of the
	//
	const (
		a = 1
		b
		c = 2
		d
	)

	fmt.Println(a, b, c, d) // "1 1 2 2"

	const Pi64 float64 = math.Pi
	var x float32 = float32(Pi64)
	var y float64 = Pi64
	var z complex128 = complex128(Pi64)

	fmt.Println(Pi64)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
