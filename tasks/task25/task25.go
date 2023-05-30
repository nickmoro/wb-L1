package main

import (
	"log"
	"time"
)

// Реализовать собственную функцию sleep.

func MySleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	log.Println("Starting:    ", time.Now())

	go func() {
		MySleep(300 * time.Millisecond)
		log.Println("From MySleep:", time.Now())
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		log.Println("From Sleep:  ", time.Now())
	}()

	time.Sleep(600 * time.Millisecond)
}
