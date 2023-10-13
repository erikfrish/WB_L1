// Задача:
// Поменять местами два числа без создания временной переменной.

// Решение:
// a, b = b, a
package main

import "fmt"

func main() {
	a := 127
	b := 42
	fmt.Printf("a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)
}
