// Задача:
// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

// Решение:
// создаем структуру Counter, содержащую sunc.RWMutex и значение value
// создаем метод для потокобезопасного инкрементирования c.Add() и получения значения c.Check()
// создаем парочку горутин для демонстрации работы счетчика
// после их выполнения выводим значение счетчика в stdout
package main

import (
	"fmt"
	"sync"
)

func main() {
	c := NewCounter()
	N := 10
	wg := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(c *Counter) {
			c.Add()
			fmt.Printf("routine %d\n", c.Check())
			wg.Done()
		}(&c)
	}
	wg.Wait()
	fmt.Printf("counter in the end is: %d\n", c.Check())
}

type Counter struct {
	value uint64
	mu    sync.RWMutex
}

func NewCounter() Counter {
	return Counter{
		mu:    sync.RWMutex{},
		value: uint64(0),
	}
}

func (c *Counter) Add() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Check() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
