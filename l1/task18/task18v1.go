package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(5)
	count := 0

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count) // Должно быть 500, т.к 5 горутин прибавляют по 100 к счетчику
}
