// Задача:
// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

// Решение:
// Код очень простой с использованием WG, все читаемо
package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	for _, num := range arr {
		wg.Add(1)
		go func(num int) {
			fmt.Println(num * num)
			wg.Done()
		}(num)
	}
	wg.Wait()
}
