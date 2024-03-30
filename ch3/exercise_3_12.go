package main

import (
	"fmt"
	"strings"
)

func isAnagram(s1, s2 string) bool {
	s1 = strings.ReplaceAll(s1, " ", "")
	s1 = strings.ToLower((s1))
	s2 = strings.ReplaceAll(s2, " ", "")
	s2 = strings.ToLower((s2))

	if len(s1) != len(s2) {
		return false
	}

	map1 := make(map[string]int)
	map2 := make(map[string]int)

	for j := 0; j < len(s1); j++ {
		map1[string(s1[j])] += 1
	}

	for j := 0; j < len(s2); j++ {
		map2[string(s2[j])] += 1
	}

	if len(map1) != len(map2) {
		return false
	}

	fmt.Print(map1, map2)

	for key, value1 := range map1 {
		value2, ok := map2[key]
		if !ok || value1 != value2 {
			return false
		}
	}

	return true
}

func main() {
	s1 := "listen"
	s2 := "silent"
	fmt.Printf("'%s' and '%s' are anagrams: %t\n", s1, s2, isAnagram(s1, s2))

	s3 := "hello"
	s4 := "world"
	fmt.Printf("'%s' and '%s' are anagrams: %t\n", s3, s4, isAnagram(s3, s4))

	s5 := "Conversation"
	s6 := "Voices Rant On"
	fmt.Printf("'%s' and '%s' are anagrams: %t\n", s5, s6, isAnagram(s5, s6))

	s8 := "Presbyterian"
	s7 := "best in prayer"
	fmt.Printf("'%s' and '%s' are anagrams: %t\n", s8, s7, isAnagram(s8, s7))
}
