package main

import "fmt"

// Программа, которая переворачивает подаваемую на ход строку. Символы могут быть unicode.
// Например: "главрыба" --> "абырвалг"

func ReverseString(str string) string {
	runes := []rune(str)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}

func main() {
	str := "главрыба"
	reversed := ReverseString(str)

	fmt.Println(str)
	fmt.Println(reversed)
}
