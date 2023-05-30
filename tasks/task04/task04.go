package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	var n int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&n)

	// Запуск воркеров, печатающих в stdout
	ch := make(chan interface{})
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		go func(index int) {
			defer wg.Done()
			for data := range ch {
				fmt.Printf("Worker %d: %v\n", index, data)
			}
		}(i)
	}

	sigsChan := make(chan os.Signal, 1)
	signal.Notify(sigsChan, syscall.SIGINT)
	i := 0
	for {
		select {
		case sig := <-sigsChan:
			// Получен сигнал
			if sig.String() == "interrupt" {
				close(ch)
				wg.Wait()
				return
			}

		default:
			// Отправляем данные разных типов
			ch <- fmt.Sprintf("Text #%d", i)
			ch <- i
			i++
		}
		time.Sleep(100 * time.Millisecond)
	}
}
