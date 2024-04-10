package main

import (
	"crypto/sha256"
	"fmt"
)

/*
When a function is called, a copy of each argument value is assigned to
the corresponding parameter variable, so the function receives a copy,
not the original.

Passing large arrays in this way can be inefficient, and any changes
that the function makes to array elements affect only the copy, not
the original.

In this regard, Go treats arrays like any other type,
but this behavior is different from languages
that implicitly pass arrays by reference.
*/
func zero(ptr *[4]int) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func main() {

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	a := &[...]int{1, 2, 3, 4}

	fmt.Printf("Before calling zero: %d\n", a)

	zero(a)
	fmt.Printf("After calling zero: %d\n", a)
}
