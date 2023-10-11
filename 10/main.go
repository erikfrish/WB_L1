// Задача:
// Дана последовательность температурных колебаний:
// -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножествах не важна.

// Решение:
// создаем функцию getGroupKey, она определяет ключ группы [key;key+10),
// в которую необходимо определить ту или иную температуру,
// потом в цикле записываем температуры в мапу по ключам,
// полученным с помощью getGroupKey,
// для демонстрации сортируем ключи и выводим в stdout значения групп
package main

import (
	"fmt"
	"sort"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	temps = append(temps, -10.0, 0.0, 10.0)

	tempGroups := make(map[int][]float64)

	getGroupKey := func(x float64) int {
		var res int
		if x > 0 {
			res = int(x/10.0) * 10
		} else if x < 0 {
			res = int(x/10.0)*10 - 10
		}
		return res
	}

	for _, temp := range temps {
		key := getGroupKey(temp)
		tempGroups[key] = append(tempGroups[key], temp)
	}

	var keys []int
	for key := range tempGroups {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		group := tempGroups[key]
		fmt.Printf("Группа (%d)-(%d) градусов: %v\n", key, key+10, group)
	}

}
