// Задача:
// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

// Решение:
// функция setBit просто меняет i-тый бит числа на указанный, как и сказано в задании

// Функция setBitOrGetNegative делает то же при i<63,
// но при изменении самого первого бита (i == 63)
// делает дополнительную проверку и
// возвращает положительное (если 0) или отрицательное (если 1) число,
// равное по модулю исходному
package main

import "fmt"

func main() {
	number := int64(16)
	fmt.Println(number)
	setBit(&number, 2, 1)
	fmt.Println(number)
	setBitOrGetNegative(&number, 63, 1)
	fmt.Println(number)
	setBitOrGetNegative(&number, 63, 0)
	fmt.Println(number)
}

func setBit(x *int64, i uint, value int) {
	if value == 1 {
		// Устанавливаем i-й бит в 1, используя побитовую операцию OR
		*x |= 1 << i
	} else if value == 0 {
		// Устанавливаем i-й бит в 0, используя побитовую операцию AND с инвертированным битом
		*x &= ^(1 << i)
	}
}

func setBitOrGetNegative(x *int64, i uint, value int) {
	if i == 63 {
		if value == 1 && *x > 0 || value == 0 && *x < 0 {
			*x = -*x
		} else {
			return
		}
	} else {
		if value == 1 {
			*x |= 1 << i
		} else if value == 0 {
			*x &= ^(1 << i)
		}
	}
}
