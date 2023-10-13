// Задача:
// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

// Решение:
// Point -- это структура, содержащая координаты (x, y),
// Line -- это структура, содержащая точки начала и конца A и B, а также длину отрезка
// x, y, length записаны с маленькой буквы, а значит доступ к ним есть только из этого пакета,
// хотя на мой взгляд их можно было бы назвать с большой буквы
// метод Line.Length() возвращает длину заданного отрезка
// SetPoint() -- конструктор Point, DrawLine -- конструктор линии, который сразу просчитывает длину
package main

import (
	"fmt"
	"math"
)

func main() {
	A := SetPoint(10, 5)
	B := SetPoint(5, 10)
	AB := DrawLine(A, B)
	fmt.Printf("A = %v, B = %v, AB = %f\n", A, B, AB.Length())
}

type Point struct {
	x, y float64
}

func SetPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

type Line struct {
	A, B   Point
	length float64
}

func distance(A, B Point) float64 {
	return math.Sqrt((math.Pow(A.x-B.x, 2)) + (math.Pow(A.y-B.y, 2)))
}

func DrawLine(A, B Point) Line {
	return Line{
		A:      A,
		B:      B,
		length: distance(A, B),
	}
}
func (l *Line) Length() float64 {
	return l.length
}
