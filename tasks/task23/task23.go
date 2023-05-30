package main

import "log"

// Удалить i-ый элемент из слайса.

// Removes element on position k in-place.
func RemoveElement(arr *[]int, pos int) {
	if pos < 0 || pos > len(*arr)-1 {
		return
	}
	*arr = append((*arr)[:pos], (*arr)[pos+1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println(arr)

	for pos := 0; pos < len(arr); pos++ {
		res := make([]int, len(arr))
		copy(res, arr)
		RemoveElement(&res, pos)
		log.Println(res)
	}

	log.Println(arr)
}
