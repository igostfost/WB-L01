/* К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

/*
	Данный код приводит к утечке памяти, так как мы создаем строку размером 1 << 10,

что равно 1024, а потом используем только 100 первых символов, а остаток остается висеть в памяти
и не освободится пока программа работает. Если мы будем вызывать много раз нашу функцию someFunc()
память неизбежно будет покидать нас)

Варианты исправления:

1. Если изменять дизайн программы нельзя, то для того чтобы освободить память, выделенную для v,
мы можем использовать конвертацию в массив байтов и переназначить этот массив nil.
Это позволит сборщику мусора освободить память.
*/
package main

import "fmt"

var justString string

func createHugeString(size int) string {
	hugeString := make([]byte, size)
	for i := 0; i < size; i++ {
		hugeString[i] = 'a'
	}
	return string(hugeString) //  - 1 вариант решения
}

func someFunc() string {
	v := createHugeString(10)

	defer func() {
		// Конвертируем v в срез байтов и переназначаем его на nil
		vBytes := []byte(v)
		for i := range vBytes {
			vBytes[i] = 0 // Опционально: обнуляем байты для безопасности
		}
		vBytes = nil // Освобождаем память, выделенную для v
	}()
	justString = v[:5]
	return justString

}

func main() {
	fmt.Println(someFunc())
}
