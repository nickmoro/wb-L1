package main

import (
	"errors"
	"fmt"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
func SetBit(num int64, bitPos, newBit int) (int64, error) {
	if bitPos < 1 || bitPos > 64 {
		return 0, errors.New("bitPos out of bounds [1; 64]")
	}
	if newBit != 0 && newBit != 1 {
		return 0, errors.New("invalid newBit: it can be 0 or 1")
	}

	bitPos--
	var mask int64 = 1 << bitPos
	if newBit == 0 {
		mask = ^mask
		return num & mask, nil
	}
	return num | mask, nil
}

func main() {
	var x int64 = 7
	fmt.Printf("%064b\n", x)

	res, _ := SetBit(x, 2, 0)
	fmt.Printf("%064b\n", res)

	res, _ = SetBit(x, 7, 1)
	fmt.Printf("%064b\n", res)
}
