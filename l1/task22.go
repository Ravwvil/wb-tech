package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int).Lsh(big.NewInt(1), 20)
	b := new(big.Int).Lsh(big.NewInt(1), 21)

	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Printf("Сумма: %s\n", sum)

	diff := new(big.Int)
	diff.Sub(a, b)
	fmt.Printf("Разность: %s\n", diff)

	product := new(big.Int)
	product.Mul(a, b)
	fmt.Printf("Произведение: %s\n", product)

	quotient := new(big.Int)
	quotient.Div(b, a)
	fmt.Printf("Частное: %s\n", quotient)
}
