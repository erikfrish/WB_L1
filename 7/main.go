// Задача:
// Реализовать конкурентную запись данных в map.

// Решение:
// я описал отдельную структуры для потокобезопасной мапы,
// содержащей саму мапу, а также мьютекс
// Реализовал специальные методы для записи и чтения данных из мапы

// в main запускаются 2 безымянные горутины, одна из которых пишет рандомные данные в канал,
// а вторая вычитывает мапу целиком и выводит ее в stdout

// можно было создать простую мапу и использовать мьютексы прямо в рутинах
// кроме того можно использовать sync.Map из стандартной библиотеки, но задача была "реализовать"
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SyncMap struct {
	m  map[any]any
	mu sync.RWMutex
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		m:  map[any]any{},
		mu: sync.RWMutex{},
	}
}

func (sm *SyncMap) Get(key any) any {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.m[key]
}
func (sm *SyncMap) GetAll() map[any]any {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	res := make(map[any]any)
	for k, v := range sm.m {
		res[k] = v
	}
	return res
}

func (sm *SyncMap) Set(key, value any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
	return
}

func main() {
	pa := NewSyncMap()

	go func(pa *SyncMap) {
		for {
			key := rand.Int()
			value := rand.Int()
			pa.Set(key, value)
			time.Sleep(time.Second * 1)
		}
	}(pa)

	go func(pa *SyncMap) {
		for {
			pam := pa.GetAll()
			fmt.Println("\nNew value of our map:", pam)
			time.Sleep(time.Second * 1)
		}
	}(pa)

	time.Sleep(time.Second * 10)
}
