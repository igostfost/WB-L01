/* Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Канал для даннных
	dataChan := make(chan string)

	// Воркер для отправки данных в канал
	go func() {
		defer close(dataChan)
		for {
			data := "myData"
			dataChan <- data
			time.Sleep(time.Second * 1)
		}
	}()

	// Воркер для чтения данных из канала
	go func(dataChan chan string) {
		for val := range dataChan {
			fmt.Println("Received data: ", val)
		}
	}(dataChan)

	// Канал для получения сигналов завершения программы
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	// Ожидаем получения сигнала о закрытии канала
	case <-signalChan:
		fmt.Println("Received signal, shutting down")

		// Если сигнал не получили, то завершаем программу через 5 секунд
	case <-time.After(time.Second * 5):
		fmt.Println("Time is left, shutting down")
	}
}
