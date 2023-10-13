// Задача:
// Реализовать все возможные способы остановки выполнения горутины.

// Решение:
// ниже в main
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// завершение рутины с помощью select и отдельного небуферизованного канала для завершения
	// а также ожидание завершения с помощью еще одного канала
	stopRoutine := make(chan bool)
	routineStopped := make(chan bool)
	go routine1(stopRoutine, routineStopped)
	time.Sleep(time.Second * 2)
	stopRoutine <- true
	<-routineStopped
	fmt.Println()

	// завершение рутины с помощью select и контекста, в данном случае с функцией отмены
	// при отмене контекста в рутине происходит return и она завершается
	ctx, cancel := context.WithCancel(context.Background())
	go routine2(ctx)
	time.Sleep(time.Second * 2)
	cancel()
	time.Sleep(time.Second)
	fmt.Println()

	// рутина сама завершится после исполнения,
	// потому что действия в ней не оформлены в бесконечный цикл
	// чего не скажешь о жизни героя
	go routine3()
	time.Sleep(time.Second * 3)
	fmt.Println()

	// завершение дочерней рутины после завершения родительской рутины, в данном случае main
	// она завершается сразу после сообщения о том, что "рабочая неделя окончена",
	// проработав ровно столько, сколько оставалось проработать main
	go routine4()
	time.Sleep(time.Second * 2)
	fmt.Println("Рабочая неделя окончена, иду бухать...")
	fmt.Println()

}

func routine1(stop <-chan bool, stopped chan<- bool) {
	for {
		select {
		default:
			fmt.Println("Занимаюсь ежедневной рутиной в свой первый рабочий день...")
			time.Sleep(1 * time.Second)
		case <-stop:
			fmt.Println("Иду домой...")
			stopped <- true
			return
		}
	}
}

func routine2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Иду в спортзал...")
			return
		default:
			fmt.Println("Занимаюсь ежедневной рутиной в свой второй рабочий день...")
			time.Sleep(1 * time.Second)
		}
	}
}

func routine3() {
	for i := 0; i < 2; i++ {
		fmt.Println("Занимаюсь ежедневной рутиной в свой третий рабочий день...")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Кем же я стал? Я хотел быть космонавтом...")
}

func routine4() {
	for {
		fmt.Println("Занимаюсь ежедневной рутиной в свой четвертый рабочий день...")
		time.Sleep(1 * time.Second)
	}
}
