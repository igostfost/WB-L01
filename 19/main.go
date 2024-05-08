// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// reverseString переворачивает строку
func reverseString(s string) string {
	// Приводим строку в срез рун, чтобы работать с юникодом
	runeStr := []rune(s)
	var builder strings.Builder

	// Идем с конца исходной строки и формируем перевернутую
	for i := utf8.RuneCountInString(s) - 1; i >= 0; i-- {
		builder.WriteRune(runeStr[i])
	}
	return builder.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Println(reverseString(scanner.Text()))
	} else {
		fmt.Println("Ошибка при считывании строки:", scanner.Err())
	}

}
