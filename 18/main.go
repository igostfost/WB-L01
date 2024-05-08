/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/
package main

import (
	"fmt"
	"sync"
)

// counter структура счетчика
type counter struct {
	value int        // значение счетчика
	mu    sync.Mutex // мютекс для конкурентной записи
}

// inc Инкрементирует значени счетчика
func (c *counter) inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func main() {
	var c counter

	workers := 100
	wg := sync.WaitGroup{}

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			c.inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.value)
}
