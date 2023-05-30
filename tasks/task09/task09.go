package main

import (
	"fmt"
	"sync"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во
// второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func Multiplier(fromArr <-chan int, toStdout chan<- int) {
	for num := range fromArr {
		toStdout <- 2 * num
	}
}

func Printer(toStdout <-chan int) {
	for num := range toStdout {
		fmt.Println(num)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fromArr, toStdout := make(chan int), make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		Multiplier(fromArr, toStdout)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Printer(toStdout)
	}()

	for _, elem := range arr {
		fromArr <- elem
	}

	close(fromArr)
	close(toStdout)

	wg.Wait()
}
