/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWords(s string) string {
	// Разбиваем на слова
	words := strings.Fields(s)
	// Переворачиваем слова в массиве с помощью индексов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	// Оюъединяем обратно в строку
	return strings.Join(words, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')
	output := reverseWords(input)
	fmt.Println(output)
}
