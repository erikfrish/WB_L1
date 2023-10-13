// Задача:
// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.

// Решение:
// Получаем от пользователя число воркеров из stdin
// Запускаем N воркеров, которые из канала в бесконечном цикле читают и выводят в stdout
// Главная горутина (главный поток) main в бесконечном цикле начинает писать рандомные данные в канал
// Программа работает бесконечно пока не будет завершена вручную
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Type number of workers:")
	var numWorkers int
	fmt.Scanln(&numWorkers)
	ch := make(chan int, numWorkers/2)

	for i := 0; i < numWorkers; i++ {
		go func(i int) {
			for {
				fmt.Println("Worker ", i, " Generated: ", <-ch)
			}
		}(i)
	}
	for {
		ch <- rand.Int()
	}
}
