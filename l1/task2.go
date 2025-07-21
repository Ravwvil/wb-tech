package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for _, digit := range arr {
		go func(digit int) {
			defer wg.Done()
			fmt.Println(digit * digit)
		}(digit)
	}
	wg.Wait()
}
