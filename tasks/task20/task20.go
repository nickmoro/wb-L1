package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: "snow dog sun" --> "sun dog snow".

func ReverseWords(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	str := "snow dog sun"
	reversedWords := ReverseWords(str)

	fmt.Println(str)
	fmt.Println(reversedWords)
}
