package main

import "fmt"

func NewStringSet(items []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set

}

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}

	set := NewStringSet(input)
	for item := range set {
		fmt.Print(item, " ")
	}
}
