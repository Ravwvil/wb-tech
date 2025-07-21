package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: stopping\n", id)
			return
		default:
			fmt.Printf("Worker %d processed job %d\n", id, i)
			i++
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// Создаём контекст, который будет отменён при получении сигнала SIGINT (Ctrl+C)
	// Этот контекст будет использоваться всеми воркерами для корректного завершения
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()

	wg := &sync.WaitGroup{}

	// Запускаем несколько воркеров, каждый из которых использует общий контекст
	// и завершится, как только контекст будет отменён (по Ctrl+C)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, wg)
	}

	fmt.Println("Press Ctrl+C to exit...")
	wg.Wait()
	fmt.Println("All workers stopped.")
}
