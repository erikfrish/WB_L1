// Задача:
// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

// Решение:
// функция reverseWords пробегает по полученной строке и записывает ее в виде [][]rune
// потом, прям как в предыдущем задании переворачиваем слайс,
// а потом перезаписываем в итоговую строку, добавляя пробелы, и возвращая ее
package main

import "fmt"

func main() {
	str := "snow dog sun"
	fmt.Printf("was: %s, is: %s\n", str, reverseWords(str))

	str = "каждый охотник желает знать где сидит фазан"
	fmt.Printf("was: %s, is: %s\n", str, reverseWords(str))
}

func reverseWords(str string) string {
	var rts [][]rune
	rts = append(rts, make([]rune, 0))
	j := 0
	for _, v := range str {
		if v == ' ' {
			rts = append(rts, make([]rune, 0))
			j++
			continue
		}
		rts[j] = append(rts[j], v)
	}
	for i := 0; i < len(rts)/2; i++ {
		rts[i], rts[len(rts)-1-i] = rts[len(rts)-1-i], rts[i]
	}
	res := ""
	for _, v := range rts[:len(rts)-1] {
		res += string(v) + " "
	}
	res += string(rts[len(rts)-1])
	return res
}
