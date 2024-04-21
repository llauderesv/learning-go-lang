package main

import "fmt"

// Go has only one mechanism to
// control the visibility of names: capitalized identifiers are exported
// from the package in which they are defined, and uncapitalized names
// are not.

type Counter struct {
	n int
}

func (c *Counter) N() int     { return c.n }
func (c *Counter) Increment() { c.n++ }
func (c *Counter) Reset()     { c.n = 0 }

func main() {
	c := Counter{0}
	c.Increment()
	c.Increment()
	c.Increment()

	fmt.Printf("The value of counter is: %d", c.N())
}
