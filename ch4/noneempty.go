package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	// Create a copy of slice string
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}

func reverse(strings []string) []string {
	out := make([]string, len(strings)) // Create a copy of slice string

	j := 0
	for i := len(strings) - 1; i >= 0; i-- {
		out[j] = strings[i]
		j++
	}

	return out
}

// func reverse2(strings []string) {
// 	pArr := &strings

// 	var tmp string
// 	j := 0
// 	for i := len(*pArr) - 1; i >= 0; i-- {
// 		tmp = (*pArr)[j] // take the first
// 		// swap the last and the first
// 		(*pArr)[j] = (*pArr)[i] // swap
// 		(*pArr)[i] = tmp

// 		fmt.Println((*pArr)[i])
// 		j++
// 	}
// }

// reverse takes a pointer to a slice of integers and reverses the slice in-place
func reverse2(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		// Swap the elements with indices i and j
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func main() {
	data := []string{"one", "", "three"}
	data2 := []string{"one", "two", "three"}

	fmt.Printf("nonempty %q\n", nonempty(data))
	fmt.Printf("nonempty2 %q\n", nonempty2(data))

	fmt.Printf("Reversed string %q\n", reverse(data2))

	num := []int{1, 2, 3}

	fmt.Printf("Before reversing num: %d\n", num)
	reverse2(&num)
	fmt.Printf("After reversing num: %d\n", num)

	// Go lang swap elements
	a := 1
	b := 2
	a, b = b, a // Swap the values of a and b

	fmt.Println(a)
	fmt.Println(b)
}
