//Реализовать бинарный поиск встроенными методами языка.
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Ввод массива с клавиатуры
	arr := make([]int, 0)
	var n int
	fmt.Print("Введите длину массива: ")
	fmt.Scan(&n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		arr = append(arr, x)
	}

	// Сортировка массива
	sort.Ints(arr)

	// Ввод элемента, который нужно найти
	var target int
	fmt.Print("Введите искомый элемент: ")
	fmt.Scan(&target)

	// Бинарный поиск элемента в отсортированном массиве
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			fmt.Printf("Элемент %d найден на позиции %d\n", target, mid)
			return
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Printf("Элемент %d не найден в массиве\n", target)
}
