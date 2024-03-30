package main

import (
	"fmt"
	"strconv"
)

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func main() {
	const GoUsage = `Go is a tool for managing Go source code.
	
	Usage:
		go command [arguments]
	`

	fmt.Println(GoUsage)

	fmt.Println(HasPrefix("Vince", "V"))
	fmt.Printf("Backing has suffix of 'ing': %v \n", HasSuffix("Backing", "ing"))

	const Name = `Vince Llauderes`

	FirstName := Name[:5]
	LastName := Name[6:]
	fmt.Println(FirstName + " " + LastName)

	for i, r := range Name {
		fmt.Printf("%d\t%q", i, r)
	}

	// Converting string to integer
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("String \"%s\" converted to integer: %d\n", str, num)

	s := Contains("Vince", "i")

	fmt.Print(s)
}
