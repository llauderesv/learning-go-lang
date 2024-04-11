package main

import "fmt"

/*
Not all panics come from the runtime. The built-in panic function maybe called directly.
A panic is often the best thing to do when some "impossible" situation happens, for instance,
execution reaches a case that logically can't happen:

Go's panic mechanism resembles exceptions in other languages
*/

func main() {
	panic(fmt.Sprint("invalid format"))
}
