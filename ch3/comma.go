package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaDecimal(s string) string {
	index := strings.LastIndex(s, ".")

	var s1, end string
	if index != -1 {
		s1 = s[:index]
		end = s[index:]
	} else {
		s1 = s
	}

	n := len(s1)
	if n <= 3 {
		return s1
	}

	return commaDecimal(s1[:n-3]) + "," + s1[n-3:] + end
}

func main() {

	fmt.Println(comma("1000000"))
	fmt.Println(commaDecimal("10000000.123"))

	s1 := "abc"
	b := []byte(s1)         // Get the byte of the string
	fmt.Print(string(b[0])) // converts to string
}
