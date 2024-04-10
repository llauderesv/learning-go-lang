package main

import (
	"fmt"
	"sort"
)

func main() {

	// Map declaration in Go
	// map[K]V where K and V are the types of its keys and values.

	ages := map[string]int{
		"charlie": 34,
		"alice":   31,
	}

	// This is equivalent to
	ages2 := make(map[string]int)
	ages2["charlie"] = 34
	ages2["alice"] = 31

	delete(ages2, "alice")

	fmt.Println(ages)
	fmt.Println(ages2["alice"]) // Ok! no error will show
	fmt.Println(ages2["charlie"])

	// Enumerate all the key/value pairs in the map, we use a range based for loop.

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// Sorting keys in Go
	var names []string
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	bob, ok := ages["bob"]
	if !ok {
		fmt.Printf("Bob is not in the Map")
	} else {
		fmt.Printf("%s is in the Map", bob)
	}

}
