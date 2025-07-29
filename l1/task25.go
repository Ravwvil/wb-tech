package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Программа стартовала:", time.Now())
	Sleep(3 * time.Second)
	fmt.Println("Программа закончила работу:", time.Now())
}
