/* Разработать программу,
которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {

	// слайс типа интерфес с набором различных типов
	variants := []any{
		21,
		"string",
		true,
		make(chan int),
	}

	for _, variant := range variants {
		// Определяем тип переменной с помощью функции TypeOf из пакета reflect
		typ := reflect.TypeOf(variant)
		fmt.Printf("Значение: %v имеет тип: %v\n", variant, typ)
	}
}
