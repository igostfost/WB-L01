/*Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов с использованием конкурентных вычислений.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var res int
	arr := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	//mu := sync.Mutex{}

	// Устанавливаем счетчик WaitGroup равным длине массива
	wg.Add(len(arr))

	for _, val := range arr {
		// Запускаем горутины
		go func(val int) {
			defer wg.Done()
			res += val * val
		}(val)
	}
	wg.Wait()

	//resultCh := make(chan int, len(arr)) // Создаем канал для передачи результата
	//
	//for _, val := range arr {
	//	wg.Add(1)
	//	go func(val int) {
	//		defer wg.Done()
	//		resultCh <- val * val // Отправляем результат в канал
	//	}(val)
	//}
	//
	//go func() {
	//	wg.Wait()
	//	close(resultCh) // Закрываем канал после завершения всех горутин
	//}()
	//
	//for r := range resultCh {
	//	res += r // Суммируем результаты из канала
	//}

	fmt.Println(res)
}
