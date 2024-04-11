package main

import "fmt"

/*
Go makes it possible to recover from a panic,
by using the recover built-in function.

A recover can stop a panic from aborting the program and let it continue with execution instead.

recover must be called within a deferred function. When the enclosing function panics, the defer will activate and a recover call within it will catch the panic.

The return value of recover is the error raised in the call to panic.
*/

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()
	fmt.Println("After mayPanic()")
}
