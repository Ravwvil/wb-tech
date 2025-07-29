package main

import "fmt"

func setBit(value int64, i uint, bit uint) int64 {
	if bit == 1 {
		return value | (1 << i)
	}
	return value &^ (1 << i)
}
func main() {
	someValue := int64(32)
	someValue = setBit(someValue, 1, 1) // Отсчет бита начинается с 0, поэтому 1 - это второй бит
	fmt.Println(someValue)
	someValue = setBit(someValue, 1, 0)
	fmt.Println(someValue)
}
