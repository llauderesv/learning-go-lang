package main

import (
	"fmt"
	"os"
)

func main() {
	// Declares a two variable s, sep with a type of string
	var s, sep string
	// The := symbol is a part of short variable declaration, a statement that declares one or more variables and gives them
	// appropriate types based on the initializer values
	// parenthesis are never used around the three components of a for loop.
	for i := 1; i < len(os.Args); i++ {
		// Concatenates the values, of the arguments in args
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// j := 1
	// for j <= 5 {
	// 	fmt.Printf("The value of j is: %d\n", j)
	// 	j += 1
	// }

	// “Go does not permit unused local variables, so this would result in a compilation error.”
	//

	// for index, arg := range os.Args[1:] {
	// 	fmt.Println(_)
	// 	s += sep + arg
	// 	sep = " "
	// }
	fmt.Println(s)
}
