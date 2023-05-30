package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.
func BinarySearch(arr []int, target int) int {
	// [left, right)
	left, right := 0, len(arr)

	for left < right {
		mid := (left + right) / 2
		if target == arr[mid] {
			return mid
		}

		if target < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for target := 1; target <= 9; target++ {
		pos := BinarySearch(arr, target)
		fmt.Println(pos, arr[pos] == target)
	}
}
