/* Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
*/

package main

import "fmt"

func main() {

	// диапозон температур
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// мапа группировки
	temperaturesGroups := map[int][]float64{}

	// группируем температуры
	for _, temp := range temperatures {
		groupsKey := int(temp/10.0) * 10
		temperaturesGroups[groupsKey] = append(temperaturesGroups[groupsKey], temp)
	}

	// выводим мапу
	fmt.Println(temperaturesGroups)

}
