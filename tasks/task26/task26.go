package main

import "fmt"

// Разработать программу, которая проверяет, что все символы в строке уникальные.
// Функция проверки должна быть регистронезависимой.

// Например:
// abcd --> true
// abCdefAaf --> false
// aabcd --> false

func IsLettersUnique(str string) bool {
	runes := []rune(str)
	set := make(map[rune]struct{}, 30)

	for _, r := range runes {
		if _, ok := set[r]; ok {
			return false
		}
		set[r] = struct{}{}
	}

	return true
}

func main() {
	str1 := "abcd"      // true
	str2 := "abCdefAaf" // false
	str3 := "aabcd"     // false

	fmt.Println(IsLettersUnique(str1), "== true")
	fmt.Println(IsLettersUnique(str2), "== false")
	fmt.Println(IsLettersUnique(str3), "== false")
}
