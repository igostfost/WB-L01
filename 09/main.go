/* Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

package main

import "fmt"

func main() {
	originalChan := make(chan int)
	doubleChan := make(chan int)

	arr := [5]int{1, 2, 3, 4, 5}

	go func() {
		defer close(originalChan)
		for _, v := range arr {
			originalChan <- v
		}
	}()

	go func() {
		defer close(doubleChan)
		for v := range originalChan {
			doubleChan <- v * 2
		}
	}()

	for data := range doubleChan {
		fmt.Println(data)
	}
}
