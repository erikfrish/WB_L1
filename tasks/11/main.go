// Задача:
// Реализовать пересечение двух неупорядоченных множеств.

// Решение:
// для нахождения пересечения двух слайсов интов используем мапу,
// если в слайсе 1 встречается значение key, то value по этому key инкрементируется
// а потом проходим по слайсу 2, если value в мапе больше нуля, то записываем в итог и отнимаем 1
package main

import "fmt"

func main() {
	arr1 := []int{3, 1, 2, 3}
	arr2 := []int{3, 6, 7, 3}

	fmt.Printf("Пересечение %v и %v: %v", arr1, arr2, Intersection(arr1, arr2))

}

func Intersection(arr1, arr2 []int) []int {
	map1 := make(map[int]int)
	result := []int{}

	for _, elem := range arr1 {
		map1[elem]++
	}

	for _, elem := range arr2 {
		if map1[elem] > 0 {
			result = append(result, elem)
			map1[elem]--
		}
	}

	return result

}
