// Задача:
// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

// var justString string
// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }
// func main() {
// 	someFunc()
// }

// Решение:
// в предоставленном коде в строчке "justString = v[:100]" есть обращение к 100 элементу, которого нет,
// решение можно реализовать двумя способами: присваивать просто v "justString = v"
// или увеличить исходную строку элементов на 100, чтобы не происходило index out of range
package main

import (
	"fmt"
	"strconv"
)

var justString string

func main() {
	someFunc1()
	someFunc2()
}
func someFunc1() {
	v := createHugeString1(1 << 10)
	justString = v
	fmt.Println(justString)
}
func someFunc2() {
	v := createHugeString2(1 << 10)
	justString = v[:100]
	fmt.Println(justString)
}
func createHugeString1(v int) string {
	return strconv.Itoa(v)
}
func createHugeString2(v int) string {
	return strconv.Itoa(v) + string(make([]byte, 100))
}
