/*Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func reverse(s string) string {
	// Создаем срез runes из рун строки s
	runes := make([]rune, utf8.RuneCountInString(s))
	// Переменная n хранит количество элементов в срезе runes
	n := len(runes)
	// Проходим по рунам в строке s и записываем их в обратном порядке в срез runes
	for _, r := range s {
		n--
		runes[n] = r
	}
	return string(runes[n:])
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')
	output := reverse(input)
	fmt.Println(output)
}
