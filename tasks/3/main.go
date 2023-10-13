// Задача:
// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

// Решение:
// Для вычисления каждого квадрата запускается отдельная горутина, они записывают данные в канал,
// из которого их читает другая горутина и считает сумму
// Для полного исполнения горутин, считающих квадраты используется WG,
// а для считающей сумму отдельный небуферизованный канал
package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	sum := 0
	sqs := make(chan int, len(arr))
	wg := sync.WaitGroup{}
	sumCalculated := make(chan bool)

	for _, num := range arr {
		wg.Add(1)
		go func(num int) {
			sqs <- num * num
			wg.Done()
		}(num)
	}

	go func() {
		for sq := range sqs {
			sum += sq
		}
		sumCalculated <- true
	}()
	wg.Wait()
	close(sqs)

	<-sumCalculated
	fmt.Println(sum)
}
