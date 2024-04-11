package main

import (
	"fmt"
)

/*
A defer statement is an ordinary function or method call prefixed by the keyword defer.
*/
func helloWorld() {
	defer fmt.Println("Vince!")
	defer fmt.Println("world") // will execute after the rest of the code below

	fmt.Println("hello")
}

// func ReadFile(filename string) ([]byte error) {
// 	f, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()
// 	return ReadAll(f)
// }

func main() {
	helloWorld()
}
