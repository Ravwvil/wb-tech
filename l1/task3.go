package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("worker %d processed job %d\n", id, j)
	}
}

func main() {
	var n int
	fmt.Print("Enter number of workers: ")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("invalid number of workers")
		return
	}

	jobs := make(chan int)
	var wg sync.WaitGroup

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	i := 0
	for {
		select {
		case jobs <- i:
			i++
		case <-sigChan:
			fmt.Println("\nShutting down...")
			close(jobs)
			wg.Wait()
			fmt.Printf("Processed %d jobs\n", i)
			return
		}
	}
}
