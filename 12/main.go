// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// создаем мапу для множества и заполняем ее в цикле элементами из множества
	sequenceMap := make(map[string]struct{})
	for _, v := range sequence {
		sequenceMap[v] = struct{}{}
	}
	fmt.Println(sequenceMap)
}
