package main

import (
	"bytes"
	"fmt"
)

func comma2(s string) string {
	var buf bytes.Buffer
	n := len(s)

	// Calculate the starting point to ensure the first group
	// has <= 3 characters but still process every 3 characters
	start := n % 3
	if start > 0 {
		buf.WriteString(s[:start])
		if n > 3 {
			buf.WriteString(",")
		}
	}

	// Iterate through the rest of the string in chunks of 3
	for i := start; i < n; i += 3 {
		buf.WriteString(s[i : i+3])
		if i+3 < n {
			buf.WriteString(",")
		}
	}

	return buf.String()
}

func main() {
	output := comma2("1231111")
	fmt.Println(output)
	// var buf bytes.Buffer

	// buf.WriteByte('v')
	// buf.WriteByte('v')
	// fmt.Fprintf(&buf, "%s", "ince")

	// buf.WriteByte('v')
	// fmt.Println(buf.String())
}
