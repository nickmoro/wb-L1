package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

// База
func GetTypeByReflect(v interface{}) string {
	return reflect.ValueOf(v).Kind().String()
}

// Мой велосипед, который кроме 4 типов ничего не поддерживает
// Причём работает только для chan из interface{}
func GetTypeBySwitch(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"

	case string:
		return "string"

	case bool:
		return "bool"

	case chan interface{}:
		return "chan interface{}"

	default:
		return "unknown type"
	}
}

func main() {
	var (
		num  int
		str  string
		flag bool
		ch   chan interface{}
	)

	fmt.Println("int:", GetTypeByReflect(num))
	fmt.Println("string:", GetTypeByReflect(str))
	fmt.Println("bool:", GetTypeByReflect(flag))
	fmt.Println("chan:", GetTypeByReflect(ch))

	fmt.Println()

	fmt.Println("int:", GetTypeBySwitch(num))
	fmt.Println("string:", GetTypeBySwitch(str))
	fmt.Println("bool:", GetTypeBySwitch(flag))
	fmt.Println("chan:", GetTypeBySwitch(ch))

}
