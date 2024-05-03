package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	dataChan := make(chan string)

	go func() {
		defer close(dataChan)
		for {
			data := "myData"
			dataChan <- data
			time.Sleep(time.Second * 1)
		}
	}()

	go func(dataChan chan string) {
		for val := range dataChan {
			fmt.Println("Received data: ", val)
		}
	}(dataChan)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signalChan:
		fmt.Println("Received signal, shutting down")
	case <-time.After(time.Second * 5):
		fmt.Println("Time is left, shutting down")
	}

}
