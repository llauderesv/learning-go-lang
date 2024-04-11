package main

import "fmt"

/*
A variadic function is one that can be called with varying numbers of arguments
To declare a variadic function the type of the final parameter is preceded by an ellipsis, "..."
*/

func sum(value ...int) int {
	total := 0
	for _, val := range value {
		total += val
	}
	return total
}

func max(vals ...int) int {
	var max int
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	min := vals[0] // start with index 0
	for _, val := range vals[1:] {
		if val < min {
			min = val
		}
	}
	return min
}

func main() {

	fmt.Println(max(100, 2, 3, 10))
	fmt.Println(max())
	fmt.Println(min(100, 1, 10))

	fmt.Println(sum(3))
	fmt.Println(sum(1, 1))
	fmt.Println(sum(1, 2, 3, 4))
}
