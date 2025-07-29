package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	left := []int{}
	right := []int{}

	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	sortedArr := quickSort(arr)
	fmt.Println(sortedArr)
}
