package main

import (
	"fmt"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечении N секунд программа должна завершаться.

func main() {
	var n int
	fmt.Print("Введите время работы программы в секундах: ")
	fmt.Scan(&n)

	// Запуск воркера, печатающего в stdout
	ch := make(chan interface{})
	wg := sync.WaitGroup{}
	go func() {
		defer wg.Done()
		for data := range ch {
			fmt.Println(data)
		}
	}()

	timerChan := time.After(time.Duration(n) * time.Second)
	i := 0
	for {
		select {
		case <-timerChan:
			// Время вышло
			close(ch)
			wg.Wait()
			return

		default:
			// Отправляем данные разных типов
			ch <- fmt.Sprintf("Text #%d", i)
			ch <- i
			i++
		}
		time.Sleep(100 * time.Millisecond)
	}
}
