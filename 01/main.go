package main

import "fmt"

/* Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	name string
	age  int
}

// Метод HumanInfo для структуры Human
func (h *Human) HumanInfo() {
	fmt.Printf("Hi, I am %s i am %d years old\n", h.name, h.age)
}

// Структура Action, в которую встроена структуры Human
type Action struct {
	Human
	Sport string
}

// Метод DoSomething для структуры Action
func (a *Action) DoSomething() {
	fmt.Printf("Hi, I am %s and I Do Something Action, my Sport: %s\n", a.name, a.Sport)
}

func main() {

	// Создаем объект Action
	action := Action{Human{"James", 20}, "Box"}

	// Вызов методов встроенной структуры Human и Action
	action.DoSomething()
	action.HumanInfo()

}
