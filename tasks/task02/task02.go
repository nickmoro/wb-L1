package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

// Конкурентность означает, что на один ресурс претендует больше пользователей,
// чем допускается этим ресурсом. В нашем случае ресурс -- срез чисел, поэтому
// каждая горутина будет работать со своим индексом массива.
func Squares(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)

	// Конечно, можно узнать число CPU (runtime.NumCPU()) и дать каждой
	// горутине по len(nums) / runtime.NumCPU() чисел. Такой подход в task02.
	// Здесь же я просто запущу len(nums) горутин, отдав распределение шедулеру.
	wg := sync.WaitGroup{}
	for i := range res {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			res[index] = res[index] * res[index]
		}(i)
	}
	wg.Wait()

	return res
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	res := Squares(nums)

	fmt.Println(nums) // [2 4 6 8 10]
	fmt.Println(res)  // [4 16 36 64 100]
}
