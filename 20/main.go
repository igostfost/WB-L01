// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWords(s string) string {
	// разделяем исходную строку по пробелам в срез слов
	words := strings.Fields(s)

	// меняем местами слова в срезе
	for l, r := 0, len(words)-1; l < r; l, r = l+1, r-1 {
		words[l], words[r] = words[r], words[l]
	}

	// формируем стркоу из слайса слов и возвращаем ее
	return strings.Join(words, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Println(reverseWords(scanner.Text()))
	} else {
		fmt.Println("Ошибка при считывании строки: ", scanner.Err())
	}
}
