package main

import (
	"fmt"
	"math"
	"time"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {

	fmt.Println(time.Now())

	fmt.Println(hypot(3, 4)) // "5"
}

// Error-Handling Strategies
/*
When a function call returns an error, it's the callers responsibility to check it and take appropriate action.

*/
