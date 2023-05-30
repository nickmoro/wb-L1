package main

import "fmt"

// Дана последовательность температурных колебаний:
// [-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5].

// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.

// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	// Будем класть значения температур в мапу слайсов.
	// Ключ в мапе получается как деление целой части температуры нацело на 10.

	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempsGrouped := make(map[int]map[float32]struct{})

	for _, temp := range temps {
		key := int(temp/10) * 10
		if _, ok := tempsGrouped[key]; !ok {
			tempsGrouped[key] = make(map[float32]struct{})
		}
		tempsGrouped[key][temp] = struct{}{}
	}

	for group, set := range tempsGrouped {
		fmt.Print(group, ": ")
		for temp := range set {
			fmt.Print(temp, " ")
		}
		fmt.Println()
	}
}
