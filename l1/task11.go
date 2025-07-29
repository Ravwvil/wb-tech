package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intersection(sliceA, sliceB []int) []int {
	maxLength := max(len(sliceA), len(sliceB))
	result := make([]int, 0, maxLength)
	set := make(map[int]struct{})
	for _, v := range sliceA {
		set[v] = struct{}{}
	}
	for _, v := range sliceB {
		if _, exists := set[v]; exists {
			result = append(result, v)
			delete(set, v)
		}
	}

	return result
}
func main() {
	sliceA := []int{1, 2, 3}
	sliceB := []int{2, 3, 4}
	result := intersection(sliceA, sliceB)
	fmt.Println(result)
}
