package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых
// переменных a, b, значение которых > 2^20.

func main() {
	// Воспользуемся "длинными" числами из пакета math/big.
	a := big.NewInt(9223372036854775807)
	b := big.NewInt(5748234928379481728)
	res := big.NewInt(0)

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("a * b =", res.Mul(a, b))
	fmt.Println("a / b =", res.Div(a, b))
	fmt.Println("a + b =", res.Add(a, b))
	fmt.Println("a - b =", res.Sub(a, b))
}
