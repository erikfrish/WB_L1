// Задача:
// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// Решение:
// быстрая сортировка работает за n*log(n) по времени.
// функция quickSort реализует быструю сортировку
// выбирается опорный элемент слайса, а потом рекурсивно сортируются все подмножества слайса,
// те, что больше и те, что меньше опорного элемента, тем же методом рекурсивно разбиваясь на подмножества
package main

import "fmt"

func main() {
	unsorted := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	quickSort(unsorted)
	fmt.Println("Отсортированный массив:", unsorted)
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Выбираем опорный элемент как середину массива
	pivot := arr[len(arr)/2]

	// Разделяем элементы на три части
	low, high := 0, len(arr)-1
	for i := low; i <= high; {
		if arr[i] < pivot {
			arr[i], arr[low] = arr[low], arr[i]
			low++
			i++
		} else if arr[i] > pivot {
			arr[i], arr[high] = arr[high], arr[i]
			high--
		} else {
			i++
		}
	}
	quickSort(arr[:low])
	quickSort(arr[high+1:])
}
