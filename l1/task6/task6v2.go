package main

import (
	"fmt"
	"time"
)

func worker(stop <-chan bool) {
	fmt.Println("Goroutine started...")
	i := 0
	for {
		select {
		default:
			fmt.Println("Current iteration:", i)
			i++
			time.Sleep(1 * time.Second)
		case <-stop:
			fmt.Println("Goroutine finished")
			return
		}
	}
}

func main() {
	// Выход из горутины через канал уведомлений
	stop := make(chan bool)
	go worker(stop)
	time.Sleep(5 * time.Second)
	stop <- true // Отправляем сигнал остановки
	fmt.Println("Program finished")
}
