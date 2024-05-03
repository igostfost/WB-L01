package main

import (
	"fmt"
	"sync"
)

/*Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			fmt.Println(val * val)
			wg.Done()
		}(val)
	}
	wg.Wait()
}
