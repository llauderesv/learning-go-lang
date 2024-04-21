package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

type Path []Point

// traditional function
// This declaration is a package-level function called geometry.Distance
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
// The extra parameter p is called method's receiver, a legacy
// from early object-oriented languages that described calling a
// method as "sending a message to an object."

// This declaration is a method of the type Point, so its name is Point.Distance
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Path is a named slice type, not a struct type like Point, yet
// we can still define methods for it.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// Because calling a function makes a copy of each argument value,
// if a function needs to update a variable, or if an argument is so
// large that we avoid copying it, we must pass the address of the variable using pointer
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	r := &Point{1, 2} // Get the pointer
	r.ScaleBy(2)
	fmt.Println(*r) // Dereference the pointer

	fmt.Println(Distance(p, q)) // "5", function call

	fmt.Println(p.Distance(q)) // "5", method call
}
