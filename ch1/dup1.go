package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	// s := "A"

	// counts[s]++
	// counts[s]++

	// counts["B"]++
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++

		if input.Text() == "exit" {
			break
		}
	}

	// NOTE: ignoring potential errors from input.Error()
	for line, n := range counts {
		// fmt.Println(n)
		if n > 0 {
			fmt.Printf("%s -> %d\n", line, n)
		}
	}
}
