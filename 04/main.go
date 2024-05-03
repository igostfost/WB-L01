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
	dataChan := make(chan string)
	var workerCount int
	fmt.Print("Input amount Workers: ")
	_, err := fmt.Scan(&workerCount)
	if err != nil {
		log.Fatal(err)
	}
	idx := 0

	go func() {
		defer close(dataChan)
		for {
			data := " data_from_main " + strconv.Itoa(idx)
			dataChan <- data
			idx++
			time.Sleep(1 * time.Second)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	for i := 0; i < workerCount; i++ {
		go func(id int, dataChan <-chan string) {
			for data := range dataChan {
				fmt.Printf("Worker %d received data: %s\n", id, data)
			}
		}(i, dataChan)
	}

	<-signalChan
}
