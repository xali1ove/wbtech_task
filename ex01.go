/*
Задание 1: Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
package main

import "fmt"

// Создаем структуру Human
type Human struct {
	Name string
	Age  int
}

// Добавим метод в структуру
func (h *Human) SayHello() {
	fmt.Printf("Привет, меня зовут %s и мне %d лет.\n", h.Name, h.Age)
}

// Встраиваем в структуру Action - структуру Human
type Action struct {
	Human
}

// Добавляем метод Walk() к структуре Action, который использует метод SayHello() из структуры Human
func (a *Action) Walk() {
	fmt.Printf("%s любит пиццу!\n", a.Name)
	a.SayHello()
}
func main() {
	// Создаем экземпляр структуры
	a := Action{
		Human{
			Name: "Руслан",
			Age:  29,
		},
	}
	// Вызываем метод на экземпляре структуры
	a.Walk()
}
