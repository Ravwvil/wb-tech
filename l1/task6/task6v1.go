package main

import (
	"fmt"
	"sync"
)

func main() {
	// Выход из горутины по условию
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("Goroutine started...")
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i == 5 {
				fmt.Println("Goroutine finished...")
				return // Выход из горутины
			}
			fmt.Println(i)
		}
	}()
	wg.Wait()
	fmt.Println("Program finished")
}
