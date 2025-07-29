package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	words := strings.Split(input, " ")
	words[len(words)-1] = strings.TrimSpace(words[len(words)-1])
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	fmt.Println(strings.Join(words, " "))
}
