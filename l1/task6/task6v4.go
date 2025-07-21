package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Остановка горутины с помощью runtime.Goexit()
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		fmt.Println("Goroutine started...")
		defer wg.Done()
		defer fmt.Println("Goroutine finished")

		fmt.Println("Goroutine started...")
		// Функция Goexit завершает горутину, в которой была вызвана
		// Goexit запускает все отложенные вызовы перед завершением горутины
		runtime.Goexit()
		fmt.Println("This will not be printed")
	}()

	wg.Wait()
	fmt.Println("Program finished")
}
