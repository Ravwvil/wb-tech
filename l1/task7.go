package main

import (
	"fmt"
	"sync"
)

func main() {
	someMap := make(map[int]int)

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			someMap[i] = i
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			someMap[1000+i] = 1000 + i
			mu.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println(someMap)
	fmt.Println(len(someMap))
}
