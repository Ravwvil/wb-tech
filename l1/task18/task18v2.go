package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)

	count := int64(0)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(count) // Должно быть 500, т.к 5 горутин прибавляют по 100 к счетчику
}
