// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
package main

import (
	"fmt"
)

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// находим опорный элемент по середине массива
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	left, right := 0, len(arr)-1

	// перемешаем опорный элемент в конец
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Проходим по всем элементам массива.
	for i := range arr {
		// Если элемент меньше опорного элемента, мы меняем его местами с левым указателем (left) и увеличиваем left
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

func main() {
	s := []int{8, 3, 5, 4, 2, 7, 6, 1, 9, 10}
	quickSort(s)
	fmt.Println(s)
}
