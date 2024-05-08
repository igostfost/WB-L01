// Реализовать паттерн «адаптер» на любом примере.
package main

import "fmt"

// интерфейс отправки сообщения
type Sender interface {
	Send(string) // Определяет метод Send для отправки сообщения
}

// Структура первого сервиса
type serviceA struct{}

// Метод Send для сервиса A
func (s *serviceA) Send(msg string) {
	fmt.Println("Send from service A", msg)
}

// Интерфейс доставки сообщения
type Deliver interface {
	Deliver(string) // Определяет метод Deliver для доставки сообщения
}

// Структура второго сервиса
type serviceB struct{}

// Метод Deliver для сервиса B
func (s *serviceB) Deliver(msg string) {
	fmt.Println("Deliver from service B", msg)
}

// Адаптер для сервиса B, чтобы он соответствовал интерфейсу Sender
type AdapterServiceB struct {
	serviceB *serviceB
}

// Метод Send адаптера для вызова метода Deliver сервиса B
func (a *AdapterServiceB) Send(msg string) {
	a.serviceB.Deliver(msg) // Вызывает метод Deliver сервиса B для отправки сообщения
}

func main() {
	A := &serviceA{}
	B := &serviceB{}

	adapterB := &AdapterServiceB{serviceB: B} // Создание адаптера для сервиса B

	sendMsg(A, "message with service A")        // Отправка сообщения с помощью сервиса A
	sendMsg(adapterB, "message with adapter B") // Отправка сообщения с помощью адаптера для сервиса B
}

// Функция отправки сообщения, принимает объект, реализующий интерфейс Sender, и отправляет сообщение через него
func sendMsg(sender Sender, msg string) {
	sender.Send(msg)
}
