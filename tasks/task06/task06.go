package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// Receiver будет остановлен по закрытию канала ch.
func Receiver(ch <-chan interface{}) {
	for data := range ch {
		log.Println("Receiver active, data =", data)
	}
	log.Println("Receiver stopping")
}

// Contexter будет остановлен по закрытию контекста через cancel-функцию.
func Contexter(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // cancel-функция вызвана извне
			log.Println("Contexter stopping")
			return

		default:
			log.Println("Contexter active")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// Spammer будет остановлен по передаче значения в канал stopCh.
func Spammer(spamCh chan<- interface{}, stopCh <-chan struct{}) {
	i := 0
	for {
		select {
		case <-stopCh:
			log.Println("Spammer stopping")
			return

		default:
			log.Println("Spammer active, data =", fmt.Sprintf("Text #%d", i))
			spamCh <- fmt.Sprintf("Text #%d", i)
			i++
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	spamCh := make(chan interface{})     // канал для общения между горутинами
	stopSpammerCh := make(chan struct{}) // канал для запроса остановки спаммера

	log.Println("All goroutines are working:")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(300)*time.Millisecond)
	defer cancel()
	go Contexter(ctx)
	go Receiver(spamCh)
	go Spammer(spamCh, stopSpammerCh)

	time.Sleep(500 * time.Millisecond)
	log.Println("Contexter already stopped")

	stopSpammerCh <- struct{}{} // остановим спаммер
	log.Println("Spammer was just stopped:")

	spamCh <- "Text from main #1"
	time.Sleep(100 * time.Millisecond)

	spamCh <- "Text from main #2"
	time.Sleep(100 * time.Millisecond)

	spamCh <- "Text from main #3"
	time.Sleep(100 * time.Millisecond)

	close(spamCh)
	time.Sleep(50 * time.Millisecond)

	log.Println("All goroutines are stopped")
	time.Sleep(200 * time.Millisecond)
}
