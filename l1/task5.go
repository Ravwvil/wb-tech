package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for job := range ch {
			fmt.Printf("worker processed job %d\n", job)
		}
	}()

	var n int
	fmt.Print("Enter number of seconds when to end the program: ")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Invalid number of seconds")
		return
	}

	timeChan := time.After(time.Duration(n) * time.Second)
	var i int
	for {
		select {
		case <-timeChan:
			fmt.Println("Timeout reached, exiting program")
			fmt.Println("Processed jobs:", i)
			close(ch)
			wg.Wait()
			return
		default:
			i++
			ch <- i
		}
	}
}
