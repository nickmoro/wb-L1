package main

import "sync"

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

// Способ 1: юзать пакет sync/atomic:
// var counter atomic.Int64 -- инициализация
// counter.Add(1) -- инкремент

// Способ 2: написать самому через простой int
type Counter struct {
	counter int
	mu      sync.RWMutex
}

func (c *Counter) Increase() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

func (c *Counter) Get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counter
}
