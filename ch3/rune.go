package main

import "fmt"

func main() {
	str := "Hello, 世界" // English "Hello," followed by Chinese characters for "world"

	// Iterating over the string to print each rune (Unicode code point)
	for _, r := range str {
		fmt.Printf("%c ", r) // %c format specifier prints the character corresponding to the rune
	}
	fmt.Println()

	// Length of the string in terms of runes
	fmt.Println("Length of the string in terms of runes:", len([]rune(str)))

	// Accessing individual characters by index
	firstRune := []rune(str)[0]
	fmt.Println("First rune: %s which equivalent to %s", firstRune, string(firstRune)) // Output: H (for English "Hello,")

	// Converting a rune to a string
	runeToStr := string(firstRune)
	fmt.Println("Rune converted to string:", runeToStr) // Output: H
}
