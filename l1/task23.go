package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный срез:", slice)
	fmt.Println("По какому индексу удалить значение?")
	var index int
	fmt.Scan(&index)
	if index < 0 || index >= len(slice) {
		fmt.Println("Индекс вне диапазона среза")
	} else {
		copy(slice[index:], slice[index+1:])
		slice = slice[:len(slice)-1] // Удаляем последний элемент, который теперь дублируется

		//slice = append(slice[:index], slice[index+1:]...) // Альтернативный способ удаления элемента по индексу
		//Возможна утечка памяти, если срез большой

		fmt.Println("Срез после удаления:", slice)
	}
}
