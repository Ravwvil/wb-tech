package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	runes := []rune(input)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	fmt.Println(string(runes))
}
