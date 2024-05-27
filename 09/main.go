/* Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

package main

import "fmt"

func main() {
	originalChan := make(chan int) // канал в который пишем числа из массива
	doubleChan := make(chan int)   // канал для результатов операции x*2

	arr := [5]int{1, 2, 3, 4, 5}

	// Воркер, который отправляет данные из массива в канал
	go func() {
		defer close(originalChan)
		for _, v := range arr {
			originalChan <- v
		}
	}()

	// Воркер который получает данные из канала originalChan и отправляет в doubleChan операцию x*2
	go func() {
		defer close(doubleChan)
		for v := range originalChan {
			doubleChan <- v * 2
		}
	}()

	// Читаем данные из канала и выводим их
	for data := range doubleChan {
		fmt.Println(data)
	}
}
