// Fahrenheit-to-Celsius conversions
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C”

	i, j := 0, 1

	fmt.Printf("The value of i=%d & the value of j=%d", i, j)
	// Excerpt From
	// The Go Programming Language (Addison-Wesley Professional Computing Series)
	// Brian W. Kernighan
	// This material may be protected by copyright.
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
