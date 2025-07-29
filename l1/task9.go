package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers []int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите числа через пробел: ")
	scanner.Scan()
	for _, number := range strings.Split(scanner.Text(), " ") {
		num, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Ошибка преобразования строки в число:", err)
			return
		}
		numbers = append(numbers, num)
	}

	FirstChan := make(chan int)
	SecondChan := make(chan int)

	go func() {
		for _, val := range numbers {
			FirstChan <- val
		}
		close(FirstChan)
	}()

	go func() {
		for val := range FirstChan {
			SecondChan <- val * val
		}
		close(SecondChan)
	}()

	for val := range SecondChan {
		fmt.Println(val)
	}
}
