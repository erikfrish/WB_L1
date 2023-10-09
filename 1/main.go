// Задача:
// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Решение:
// Встраиваем Human в Action, включая все методы. Функция main демонстрирует, что это работает
package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	isMale bool
}

func (h *Human) Born(name string, isMale bool) {
	h.Name = name
	h.Age = 0
	h.isMale = isMale
}

func (h *Human) ChangeSex() {
	h.isMale = !h.isMale
}

func (h *Human) ChangeName(name string) {
	h.Name = name
}

func (h *Human) HowOld() int {
	return h.Age
}

func (h *Human) BecomeOlderByAYear() {
	h.Age++
}

func (h *Human) BecomeOlderByNYeras(n int) {
	h.Age += n
}

type Action struct {
	Human
}

func main() {
	var le_ha Action
	le_ha.Born("Le_ha", true)
	fmt.Println(le_ha)
	le_ha.BecomeOlderByAYear()
	fmt.Println(le_ha)
	le_ha.BecomeOlderByNYeras(18)
	le_ha.ChangeSex()
	fmt.Println(le_ha)
}
