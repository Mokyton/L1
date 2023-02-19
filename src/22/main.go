package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	a := big.NewInt(int64(math.Pow(2, 20)))
	b := big.NewInt(int64(math.Pow(2, 20)))

	sum := new(big.Int).Add(a, b)
	fmt.Println("a + b =", sum)
	sub := new(big.Int).Sub(a, b)
	fmt.Println("a - b =", sub)
	mul := new(big.Int).Mul(a, b)
	fmt.Println("a * b =", mul)
	div := new(big.Int).Div(a, b)
	fmt.Println("a / b =", div)
}
