// 1. Дана структура Human с произвольным набором полей и методов. Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования)

package main

import "fmt"

// Создаём структуру Human
type Human struct {
	name string
	age  int
}

// Метод Human
func (h *Human) Greetings() {
	fmt.Printf("Привет! Меня зовут %s, и мне %d лет\n", h.name, h.age)
}

type Action struct {
	// встраиваем анонимное поле
	Human
}

// Метод Action
func (a *Action) Drink() {
	fmt.Printf("Ну вот, %s опять пошёл в бар вместо работы\n", a.name)
}

func main() {
	someHuman := Human{
		name: "Саня",
		age:  25,
	}
	action := Action{Human: someHuman}
	// можем вызвать у action оба метода
	action.Greetings()
	action.Drink()
}
