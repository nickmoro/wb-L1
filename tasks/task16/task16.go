package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func split(arr *[]int, left, right int) int {
	// Берём самый правый элемент массива как pivot.
	// Слева от pivot элементы, меньшие его.
	// Справа от pivot элементы, большие его.

	pivot := (*arr)[right]
	mid := left
	for j := left; j < right; j++ {
		if (*arr)[j] < pivot {
			(*arr)[mid], (*arr)[j] = (*arr)[j], (*arr)[mid]
			mid++
		}
	}
	(*arr)[mid], (*arr)[right] = (*arr)[right], (*arr)[mid]
	return mid
}

func quickSort(arr *[]int, left, right int) {
	if left < right {
		mid := split(arr, left, right)
		quickSort(arr, left, mid-1)
		quickSort(arr, mid+1, right)
	}
}

func main() {
	n := 50

	rand.Seed(time.Now().UnixMicro())
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("Unsorted:", arr)

	// Самописная сортировка
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	quickSort(&sorted, 0, len(arr)-1)

	fmt.Println("Sorted:  ", sorted)

	// Сортировка из пакета sort
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	fmt.Println("Sorted:  ", arr)
}
