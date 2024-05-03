// Реализовать возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func withChannel(closeChannel chan struct{}) {
	fmt.Println("Method 1: Communication channels")
	for {
		select {
		case <-closeChannel:
			fmt.Println("Method 1: Stop")
			return
		default:
			fmt.Println("Method 1: Working")
		}
		time.Sleep(time.Second)
	}
}

func withContext(ctx context.Context) {
	fmt.Println("Method 2: Communication context")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Method 2: Stop")
			return
		default:
			fmt.Println("Method 2: Working")
		}
		time.Sleep(time.Second)
	}
}

func receiver(dataChan <-chan int) {
	for {
		data, ok := <-dataChan
		if !ok {
			fmt.Println("Method 5: Channel closed")
			return
		}
		fmt.Println("Method 5: Working -", data)
	}

}

func sender(dataChan chan<- int) {
	for i := 1; i < 3; i++ {
		dataChan <- i
	}
}

func main() {

	// 1. Использование каналов для коммуникации
	stopChan := make(chan struct{})
	go withChannel(stopChan)

	time.Sleep(time.Second * 2)
	stopChan <- struct{}{}
	close(stopChan)

	// 2. Использование контекстов
	ctx, cancel := context.WithCancel(context.Background())
	go withContext(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(time.Second * 1)

	ctxTimer, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2()
	go withContext(ctxTimer)
	time.Sleep(time.Second * 2)

	// 3. Использование таймеров или таймаутов
	timeout := time.After(time.Second * 2)
	go func() {
		fmt.Println("Method 3: Timers or timeouts")
		for {
			select {
			case <-timeout:
				fmt.Println("Method 3: Stop")
				return
			default:
				fmt.Println("Method 3: Working")
			}
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 3)

	// 4. Принудительная остановка
	go func() {
		fmt.Println("Method 4: Forced")
		for {
			fmt.Println("Method 4: Working")
			runtime.Goexit()
		}
	}()
	time.Sleep(time.Second * 2)

	// 5. При попытки чтения из закрытого канала
	chan5 := make(chan int)
	go sender(chan5)
	go receiver(chan5)
	time.Sleep(time.Second * 2)
	close(chan5)
	time.Sleep(time.Second * 1)
}
