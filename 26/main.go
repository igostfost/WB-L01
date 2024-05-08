/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isUnique(s string) bool {
	// меняем регистр на строчные символы
	s = strings.ToLower(s)

	// создаем мапу для хранения символов
	charMap := make(map[rune]bool)

	// проходим по строке, заполняем мапу символами
	for _, c := range s {
		// Если такой символ уже есть, возвращаем false
		if charMap[c] {
			return false
		}
		charMap[c] = true
	}

	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(isUnique(scanner.Text()))
}
