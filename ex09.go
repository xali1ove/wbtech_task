/*Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

package main

import "fmt"

func FirstToSecond(f, s chan int) {
	for val := range f {
		s <- val * 2
	}
	close(s)
}

func main() {
	firstCh := make(chan int)
	secondCh := make(chan int)

	go FirstToSecond(firstCh, secondCh)

	var n int
	fmt.Println("Enter the number of values:")
	fmt.Scanln(&n)
	inputVals := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter value #%d: ", i+1)
		fmt.Scanln(&inputVals[i])
	}

	// Отправляем значения в первый канал
	go func() {
		defer close(firstCh)
		for _, val := range inputVals {
			firstCh <- val
		}
	}()

	fmt.Println("Result:")
	for val := range secondCh {
		fmt.Println(val)
	}
}
