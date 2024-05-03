// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0

package main

import "fmt"

func setBit(val *int64, pos int, bitVal int) {
	if bitVal == 1 {
		*val = *val | 1<<pos // 0010 = 0010 | 1000 = 1010
	} else if bitVal == 0 {
		*val = *val & ^(1 << pos) // 0010 = 0010 & ^(0100) = 0010 & 1011 = 0010
	} else {
		fmt.Println("Invalid Bit Value. Only 0 or 1")
	}
}
func main() {

	var x int64 = 2
	fmt.Printf("Part 1: binary: %05b, dec: %d\n", x, x)
	setBit(&x, 3, 1)
	fmt.Printf("Part 2: binary: %05b, dec: %d\n", x, x)
}
