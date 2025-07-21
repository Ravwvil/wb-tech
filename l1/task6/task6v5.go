package main

import (
	"fmt"
	"sync"
)

func worker(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine started...")
	for data := range ch {
		fmt.Printf("Received: %d\n", data)
	}
	fmt.Println("Channel closed. Goroutine finished.")
}

func main() {
	// Остановка горутины через закрытие канала (чтение из канала с помощью range)
	dataChannel := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go worker(dataChannel, wg)

	for i := 1; i <= 5; i++ {
		dataChannel <- i
	}

	close(dataChannel)
	wg.Wait()
	fmt.Println("Main function finished.")
}
