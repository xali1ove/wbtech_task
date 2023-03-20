/*Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/
package main

import (
	"fmt"
	"time"
)

// функция отправки данных в канал
func sendData(ch chan<- int) {
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
}

// функция чтения данных из канала
func readData(ch <-chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	var timer int
	fmt.Println("Введите количество секунд: ")
	fmt.Scanln(&timer)

	ch := make(chan int)

	go sendData(ch) // запускаем отправку данных в канал
	go readData(ch) // запускаем чтение данных из канала

	// ожидание завершения работы программы
	time.Sleep(time.Duration(timer) * time.Second)
	fmt.Println("Программа завершена.")
}
