//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	// создаем пустое множество
	wordSet := make(map[string]bool)

	// добавляем каждое уникальное слово в множество
	for _, word := range words {
		wordSet[word] = true
	}

	// выводим элементы множества
	for word := range wordSet {
		fmt.Println(word)
	}
}
