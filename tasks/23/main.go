// Задача:
// Удалить i-ый элемент из слайса.

// Решение:
// функция delete удаляет i-тый элемент слайса и возвращает итоговое его значение.
// так как под капотом используется append, исходный массив изменяется, поэтому
// использовать delete необходимо также как и append, присваивая исходному слайсу результат:
// arr = delete(arr, i)
// в main заводим слайс и удаляем его элемент по индексу 1
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("arr before delete: %v\n", arr)
	arr = delete(arr, 1)
	fmt.Printf("arr after delete: %v\n", arr)
}

func delete(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}
