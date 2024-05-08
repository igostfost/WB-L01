// Удалить i-ый элемент из слайса.

package main

import "fmt"

func RemoveElemFromSlice(nums []int, pos int) []int {
	return append(nums[:pos], nums[pos+1:]...)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(RemoveElemFromSlice(s, 1))
}
