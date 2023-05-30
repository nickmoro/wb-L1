package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
type Human struct {
	Name string
	Age  int
}

func (h *Human) Introduce() string {
	return fmt.Sprintf("I am %s", h.Name)
}

// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).
type Action struct {
	Human // доступны все поля и методы структуры Human
}

func (a *Action) Act() string {
	introStr := a.Introduce() // we can use Introduce from Human
	return fmt.Sprintf("%d years old said: %s", a.Age, introStr)
}

func main() {
	a := Action{Human{Name: "Ivan", Age: 21}}

	fmt.Println(a.Name)        // доступны поля структуры Human
	fmt.Println(a.Age)         // доступны поля структуры Human
	fmt.Println(a.Introduce()) // доступны методы структуры Human

	// Доступен метод Act, использующий метод Human.Introduce под капотом
	fmt.Println(a.Act())
}
