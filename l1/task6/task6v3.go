package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Остановка горутины с помощью контекста
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		fmt.Println("Goroutine started...")
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context cancelled. Goroutine finished") // Контекст отменен, завершаем работу
				return
			default:
				fmt.Println(i)
				i++
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)
	wg.Wait()
	fmt.Println("Program finished")
}
