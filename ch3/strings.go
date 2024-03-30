package main

// The strconv package provides functions for converting boolean,
// integer, and floating-point values to and from their representation
// and functions for quoting and unquoting strings.

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(strconv.Atoi("123"))

	pie := "3.14"

	realPie, err := strconv.ParseFloat(pie, 32)
	if err != nil {
		fmt.Printf("Error converting pie to float: %s", err)
	}

	fmt.Printf("Pie value is: %f", realPie)

}
