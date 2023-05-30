package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Дана последовательность чисел: {2, 4, 6, 8, 10}. Найти сумму их квадратов
// (2^2 + 3^2 + 4^2 + ...) с использованием конкурентных вычислений.
func SquaresSum(nums []int) int {
	return Sum(Squares(nums)) // Квадраты считаем при помощи функции из task02
}

// Sum возвращает сумму элементов массива
func Sum(nums []int) int {
	// Каждая горутина суммирует равное количество чисел, но не менее двух
	// Исключение составляет последняя горутина, если len(nums) % numCPU != 0
	forWorker := len(nums) / runtime.GOMAXPROCS(0)
	if forWorker < 2 {
		forWorker = 2
	}

	workersNum := len(nums) / forWorker // Число необходимых горутин
	res := make([]int, workersNum)      // Каждая горутина запишет результат в свой индекс

	wg := sync.WaitGroup{}
	left := 0
	for i := 0; i < workersNum; i++ {
		right := left + forWorker
		if right >= len(nums) {
			right = len(nums) - 1
		}

		wg.Add(1)
		go func(index, left, right int) {
			// [left; right]
			defer wg.Done()
			sum := 0
			for j := left; j <= right; j++ {
				sum += nums[j]
			}
			res[index] = sum
		}(i, left, right)

		left = right + 1
	}

	wg.Wait()

	if len(res) == 1 {
		return res[0]
	}

	return Sum(res)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	res := SquaresSum(nums)

	fmt.Println(nums) // [2 4 6 8 10]
	fmt.Println(res)  // 220
}

// Из task02
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
