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
	mu := sync.Mutex{}
	wg.Add(len(arr))
	for _, val := range arr {
		go func(val int, res *int) {
			mu.Lock()
			*res += val * val
			mu.Unlock()
			defer wg.Done()
		}(val, &res)
	}
	wg.Wait()
	fmt.Println(res)
}
