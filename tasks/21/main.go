// Задача:
// Реализовать паттерн «адаптер» на любом примере.

// Решение:
// жизненный пример адаптера -- это адаптер с квадратной вилки США а знакомую нам круглую европейскую
// EUSocket -- это европейская розетка, к ней можно подключить только европейскую вилку,
// это видно по сигнатуре метода connect (s *EUSocket) connect(plug EUPlug)
// чтобы подключить зарядку myMacBookCharger, имеющую US вилку используется адаптер UStoEUAdapter{}
// создаем myAdapter, вешаем его на зарядку от macbook -- myAdapter.putOnPlug(&myMacBookCharger)
// а затем подключаем к розетке -- myBedroomSocket.connect(&myAdapter)
package main

import "fmt"

func main() {
	myBedroomSocket := EUSocket{color: "beige"}
	myMacBookCharger := USCharger{shape: "square"}
	myAdapter := UStoEUAdapter{}
	myAdapter.putOnPlug(&myMacBookCharger)
	myBedroomSocket.connect(&myAdapter)
}

type EUPlug interface {
	connectToEUSocket()
}

type USPlug interface {
	connectToUSSocket()
}

type USCharger struct {
	shape string
}

func (c *USCharger) connectToUSSocket() {
	fmt.Println("US charger connected")
}

type EUSocket struct {
	color string
}

func (s *EUSocket) connect(plug EUPlug) {
	plug.connectToEUSocket()
}

type UStoEUAdapter struct {
	device USPlug
}

func (a *UStoEUAdapter) putOnPlug(device USPlug) {
	a.device = device
}

func (a *UStoEUAdapter) connectToEUSocket() {
	a.device.connectToUSSocket()
	fmt.Println("to EU socket via adapter")
}
