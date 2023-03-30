/*Дана переменная int64.
Разработать программу которая устанавливает i-й бит в 1 или 0.*/
package main

import (
	"fmt"
)

func setBit(num int64, i uint, value int) int64 {
	// Устанавливаем i-й бит в заданное значение
	var result int64
	if value == 0 {
		result = num &^ (1 << i)
	} else if value == 1 {
		result = num | (1 << i)
	} else {
		panic("Неправильный ввод")
	}

	return result
}

func main() {

	var num int64
	var i uint
	var ZeroOne int

	fmt.Print("Введите число: ")
	fmt.Scanln(&num)

	fmt.Print("Введите номер бита: ")
	fmt.Scanln(&i)

	fmt.Print("Введите 0 или 1: ")
	fmt.Scanln(&ZeroOne)

	// Устанавливаем i-й бит в заданное значение
	result := setBit(num, i, ZeroOne)

	fmt.Printf("Результат: %d\n", result)
}
