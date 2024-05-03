package main

import "fmt"

/* Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	name string
	age  int
}

type Student struct {
	Name  string
	Point int
}

func (s *Student) teach() {
	fmt.Printf("Hi i am student -  %s, and my pooint: %d\n", s.Name, s.Point)
}

func (h *Human) personInfo() {
	fmt.Printf("Hi, I am %s i am %d years old\n", h.name, h.age)
}

type Action struct {
	Human
	Student
	Sport string
}

func main() {
	alex := Action{
		Sport: "Run",
		Human: Human{
			name: "Alex",
			age:  25},
	}

	alex.personInfo()
	fmt.Printf("Human name: %s, Human sport: %s\n", alex.name, alex.Sport)

	Ivan := Action{
		Sport:   "GOGO",
		Student: Student{Name: "Ivan", Point: 20},
	}

	Ivan.teach()
	fmt.Printf("Student name: %s, Student sport: %s\n", Ivan.Name, Ivan.Sport)
}
