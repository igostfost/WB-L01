// Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"strconv"
	"sync"
)

type SafeMap struct {
	sync.Mutex
	m map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]int)}
}

func (s *SafeMap) Set(key string, value int) {
	s.Lock()
	s.m[key] = value
	s.Unlock()
}

func main() {

	myMap := NewSafeMap()
	wg := sync.WaitGroup{}
	wg.Add(10)

	// Записываем конкуренто данные в мапу
	go func() {
		for i := 0; i < 10; i++ {
			myMap.Set(strconv.Itoa(i), i*10)
			wg.Done()
		}
	}()
	wg.Wait()
	fmt.Println(myMap.m)
}

// Безопасную запись обеспечивает встроенный в структуру мапы мютекс
