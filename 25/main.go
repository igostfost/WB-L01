// Реализовать собственную функцию sleep.
package main

import (
	"fmt"
	"time"
)

// Sleep задержка на указанное количество секунд
func Sleep(duration int) {
	<-time.After(time.Second * time.Duration(duration))
}

func main() {

	start := time.Now()
	fmt.Println(start.Format(time.Stamp), "sleeping to 3 sec ...")
	Sleep(3)
	end := time.Now()
	fmt.Println(end.Format(time.Stamp))
}
