// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func findSetIntersect(s1, s2 []int) []int {
	// слайс для результата
	intersect := make([]int, 0)

	// создаем мапу для первого множества и заполняем его данными из первого множества
	s1map := make(map[int]struct{})
	for _, v := range s1 {
		s1map[v] = struct{}{}
	}

	// проходим по второму множеству и если в мапе первого множества есть ключ из второго множества - это пересечение
	for _, v := range s2 {
		if _, ok := s1map[v]; ok {
			intersect = append(intersect, v)
		}
	}

	return intersect
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 7, 2, 9, 3}
	fmt.Println(findSetIntersect(s1, s2))
}
