/*Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	// // Создаем канал для передачи данных
	dataChan := make(chan string)

	// Запрашиваем количество воркеров
	var workerCount int
	fmt.Print("Input amount Workers: ")
	_, err := fmt.Scan(&workerCount)
	if err != nil {
		log.Fatal(err)
	}
	idx := 0

	// Воркер для генерации данных и отправки их в канал
	go func() {
		defer close(dataChan) // Закрываем канал после завершения генерации данных
		for {
			data := " data_from_main " + strconv.Itoa(idx)
			dataChan <- data
			idx++
			time.Sleep(1 * time.Second)
		}
	}()

	// Канал для получения сигналов завершения программы
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Создаем указанное количесвто воркеров
	for i := 0; i < workerCount; i++ {
		go func(id int, dataChan <-chan string) {

			// Читаем данные из канала и обрабатываем их в каждом воркере
			for data := range dataChan {
				fmt.Printf("Worker %d received data: %s\n", id, data)
			}
		}(i, dataChan)
	}

	// Ожидаем сигнала завершения программы
	<-signalChan
}
