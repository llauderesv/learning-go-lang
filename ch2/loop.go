package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "watermelon"}
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}
}
