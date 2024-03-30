package main

import (
	"fmt"
	"os"
)

func main() {

	var s, sep string
	for index, arg := range os.Args[1:] {
		fmt.Println(index)
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
