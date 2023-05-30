package main

import "strings"

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
/*
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
*/

// Судя по всему создаётся строка из 2^10 символов, которая потом обрезается до 100.
// Соответственно, кажется логичным просто сразу создать строку из 100 символов.

// Помимо этого, строка создаётся как глобальная переменная. Конечно, на стеке будет лежать
// только хедер строки, а сами байты будут в куче, но всё равно использование глобальной
// переменной должно быть обосновано.

// Исправленный код:

func someFunc1(str *string) {
	*str = createHugeString(100)
}

func someFunc2() string {
	return createHugeString(100)
}

func main() {
	// Вариант 1, если по каким-то причинам необходимо объявить строку до вызова функции
	var justString string
	someFunc1(&justString)

	// Вариант 2, если по каким-то причинам надо обернуть вызов createHugeString в someFunc
	justString = someFunc2()

	// Вариант 3, самый адекватный
	justString = createHugeString(100)

	// P.S.: Возможно, мы захотим передавать n int в качестве параметра в someFunc.
	// 	Для финального выбора нужно смотреть на данный абстрактный проект/код целиком.
}

func createHugeString(n int) string {
	return strings.Repeat("a", n)
}
