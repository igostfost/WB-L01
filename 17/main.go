// Реализовать бинарный поиск встроенными методами языка.
package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	if right < left {
		return -1
	}

	for left <= right {
		mid := (left + right) / 2 // Находим индекс центра
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target { // Если искомый элемент  больше центрального элемента
			left = mid + 1 // Значит он находится справа и двигаемся в право, отсекая левую часть
		} else {
			right = mid - 1
		}
	}
	return -1 // Элемент не найден
}

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := binarySearch(s, 5)
	if index != -1 {
		fmt.Printf("Элемент найден в позиции %d\n", index)
	} else {
		fmt.Printf("Элементне найден в массиве\n")
	}

}
