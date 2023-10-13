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
// Завершение происходит при получении сигнала от ОС.
// Если бы данные писались из main, тогда каждый воркер бы ожидал,
// пока сам main запишет данные в канал, а значит остановка main без всяких
// дополнительных контекстов, каналов и горутин завершала бы все воркеры
// как дочерние горутины,
// но для явной обработки сигнала и реализации принципа graceful shutdown
// используется отдельная горутина, которая ожидает сигнала о завершении и после
// отменяет контекст, на который ориентируется горутина записывающая данные в канал
// и все воркеры
package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("Type number of workers:")
	var numWorkers int
	fmt.Scanln(&numWorkers)
	ch := make(chan int, numWorkers/2)

	ctx, cancel := context.WithCancel(context.Background())

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)

	for i := 0; i < numWorkers; i++ {
		go func(ctx context.Context, i int) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Println("Worker ", i, " Generated: ", <-ch)
				}
			}
		}(ctx, i)
	}
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- rand.Int()
			}
		}
	}(ctx)
	go func() {
		for range signalChan {
			fmt.Println("\nReceived an interrupt, exiting")
			cancel()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
