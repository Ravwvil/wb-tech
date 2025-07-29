package main

import "fmt"

func detectType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Тип: int, значение:", v)
	case string:
		fmt.Println("Тип: string, значение:", v)
	case bool:
		fmt.Println("Тип: bool, значение:", v)
	case chan int:
		fmt.Println("Тип: chan int")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	var a int = 19
	var b string = "Hello"
	var c bool = true
	var d chan int = make(chan int)

	detectType(a)
	detectType(b)
	detectType(c)
	detectType(d)
	detectType(3.22)
}
