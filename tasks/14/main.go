// Задача:
// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}

// Решение:
// для определения типа переменной в go удобнее всего использовать встроенную библиотеку reflect
// и функцию TypeOf оттуда, она отлично справляется
// предположу, что это не достаточно полное решение, поэтому написал кастомную функцию customTypeOf,
// которая определяет тип с помощью switch v.(type), тем более раз задан список типов, которые нужно определять
// в моей реализации определяется только chan any, но при желании можно добавить проверку на канал с любым типом
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var va any
	va = 0.0
	fmt.Println("va = ", va)
	fmt.Println("reflect.TypeOf of va =", reflect.TypeOf(va))

	va = "some text"
	fmt.Println("va = ", va)
	fmt.Println("reflect.TypeOf of va =", reflect.TypeOf(va))

	va = make(chan any)
	fmt.Println("va = ", va)
	fmt.Println("customTypeOf of va =", customTypeOf(va))

	va = false
	fmt.Println("va = ", va)
	fmt.Println("customTypeOf of va =", customTypeOf(va))
}

func customTypeOf(va any) string {
	switch va.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan any:
		return "chan"
	default:
		return "smth"
	}
}
