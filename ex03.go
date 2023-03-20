/*Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42....) с использованием конкурентных вычислений.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	//для хранения суммы квадратов чисел
	sum := 0
	//для ожидания завершения всех горутин
	var wg sync.WaitGroup
	//добавляем количество горутин, которые будут использоваться для обработки чисел
	wg.Add(len(numbers))
	for _, n := range numbers {
		//запускаем горутину для каждого числа в последовательности
		go func(x int) {
			//уведомление что рутина завершена
			defer wg.Done()
			sum += x * x
		}(n)
	}
	//дождаться завершения всех горутин
	wg.Wait()
	fmt.Printf("Сумма квадратов чисел %v равна %d\n", numbers, sum)
}