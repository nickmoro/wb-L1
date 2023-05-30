package main

import (
	"fmt"
	"sync"
)

// Реализовать паттерн «адаптер» на любом примере.

// Класс Adapter приводит интерфейс класса Adaptee в соответствие с интерфейсом класса Target,
// 	который реализуется классом Adapter.
// Это позволяет объекту Client использовать объект Adaptee (посредством адаптера) так,
// 	словно он является экземпляром класса Target.

type MutexAdapter struct {
	mu sync.Mutex
}

func (a *MutexAdapter) Close() {
	a.mu.Lock()
}

func (a *MutexAdapter) Open() {
	a.mu.Unlock()
}

type Mutex interface {
	Close()
	Open()
}

func main() {
	// Исторически сложилось, что программа во многих местах использует функцию mutex.Close().
	// По некоторым причинам, теперь необходимо использовать метод mutex.Lock() из другого пакета.
	// Чтобы не править проект во многих местах, был реализован адаптер, реализующий интерфейс.

	mu := MutexAdapter{} // реализует интерфейс Mutex

	mu.Close() // можно использовать Close()
	mu.Open()  // можно использовать Open()

	fmt.Println("OK")
}
