//Реализовать быструю сортировку массива (quicksort) встроенными методами языка
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Чтение чисел из стандартного ввода
	fmt.Print("Введите числа через пробел: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Разбиение введенной строки на отдельные числа
	numStrings := strings.Fields(input)
	nums := make([]int, len(numStrings))

	for i, nStr := range numStrings {
		num, err := strconv.Atoi(nStr)
		if err != nil {
			fmt.Println("Ошибка: введено не число")
			return
		}
		nums[i] = num
	}

	// Сортировка и вывод отсортированного массива
	sort.Ints(nums)
	fmt.Println("Отсортированный массив:", nums)
}
