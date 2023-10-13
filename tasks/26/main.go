// Задача:
// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

// Решение:
// функция isUnique(string) возвращает true, если символы в строке уникальные,
// и false, если неуникальные, она не зависит от регистра,
// так как первым делом все символы в строке приводятся к строчному виду
package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "abcd"
	fmt.Printf("str: %s, isUnique(str): %t\n", str1, isUnique(str1))
	str2 := "abCdefAaf"
	fmt.Printf("str: %s, isUnique(str): %t\n", str2, isUnique(str2))
	str3 := "aabcd"
	fmt.Printf("str: %s, isUnique(str): %t\n", str3, isUnique(str3))
}

func isUnique(str string) bool {
	str = strings.ToLower(str)
	ma := make(map[rune]uint16, len(str))
	for _, letter := range str {
		ma[letter]++
		if ma[letter] > 1 {
			return false
		}
	}
	return true
}
