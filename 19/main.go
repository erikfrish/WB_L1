// Задача:
// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

// Решение:
// для переворачивания строки описана отдельная функция reverse(str string) string,
// она создает []rune, чтобы итерироваться посимвольно, а не побайтно, ведь кириллица занимает 2 байта,
// а потом в цикле меняет местами элементы попарно (первый-последний, второй-предпоследний и т.д.)
// таким образом не задействуется лишняя память и в цикле мы проходим всего len(str)/2 итераций
package main

import "fmt"

func main() {
	str := "главрыба"
	fmt.Printf("was: %s, is: %s\n", str, reverse(str))

	str = "птицемясокомбинат"
	fmt.Printf("was: %s, is: %s\n", str, reverse(str))

	str = "бу"
	fmt.Printf("was: %s, is: %s\n", str, reverse(str))

	str = "у"
	fmt.Printf("was: %s, is: %s\n", str, reverse(str))
}

func reverse(str string) string {
	rts := []rune(str)
	for i := 0; i < len(rts)/2; i++ {
		rts[i], rts[len(rts)-1-i] = rts[len(rts)-1-i], rts[i]
	}
	return string(rts)
}
