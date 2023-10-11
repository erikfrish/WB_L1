// Задача:
// Реализовать бинарный поиск встроенными методами языка.

// Решение:
// функция search выполняет бинарный поиск по отсортированному слайсу и возвращает индекс найденного элемента
// поиск осуществляется с использованием двух указателей left и right,
// с каждой итерацией один из указателей пододвигается и сужает область поиска в 2 раза,
// таким образом для поиска элемента требуется log2(n) итераций
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	x := 10
	sort.Ints(arr)
	fmt.Printf("Index of %d element is %d\n", x, search(arr, x))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	p := len(nums) / 2
	for {
		if nums[p] == target {
			break
		}
		if left == p {
			if nums[right] == target {
				p = right
			} else {
				p = -1
			}
			break
		}
		if target < nums[p] {
			right = p
			p = left + (right-left)/2
		}
		if target > nums[p] {
			left = p
			p = left + (right-left)/2
		}
	}
	return p
}
