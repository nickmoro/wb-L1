package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func XorSwap(a, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}

func main() {
	a, b := 3, 4

	swappedA1, swappedB1 := b, a // 1 вариант

	swappedA2, swappedB2 := a, b
	XorSwap(&swappedA2, &swappedB2) // 2 вариант

	fmt.Println(a, b)
	fmt.Println(swappedA1, swappedB1)
	fmt.Println(swappedA2, swappedB2)
}
