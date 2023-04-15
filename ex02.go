/*
Задание 2: Написать программу на Go, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	//создаем канал для передачи вычисленных квадратов чисел
	squareCh := make(chan int)
	//создаем WaitGroup wg, чтобы дождаться завершения всех вычислений
	var wg sync.WaitGroup
	wg.Add(len(nums))
	//запускаем горутину для каждого числа
	for _, n := range nums {
		go func(x int) {
			defer wg.Done()
			squareCh <- x * x
		}(n)
	}
	//горутины завершили вычисления, мы закрываем канал squareCh и переходим к чтению из канала
	go func() {
		wg.Wait()
		close(squareCh)
	}()
	//печатаем каждое значение, полученное из канала, в stdout
	for square := range squareCh {
		fmt.Println(square)
	}
}
