package main

import (
	"fmt"
	"strings"
)

func hasUniqueChars(s string) bool {
	s = strings.ToLower(s)
	seen := make(map[rune]bool)

	for _, ch := range s {
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}

func main() {
	fmt.Println(hasUniqueChars("abcd"))      // true
	fmt.Println(hasUniqueChars("abCdefAaf")) // false
	fmt.Println(hasUniqueChars("aabcd"))     // false
}
