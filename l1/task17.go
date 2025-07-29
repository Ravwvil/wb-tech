package main

import "fmt"

func binSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func main() {
	arr := []int{1, 2, 2, 3, 5, 6, 6, 6, 8, 9, 10}
	target := 3
	result := binSearch(arr, target)
	if result != -1 {
		fmt.Println("Элемент найден по индексу:", result)
	} else {
		fmt.Println("Элемент не найден")
	}
}
