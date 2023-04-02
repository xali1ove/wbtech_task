//Возвращаем значение через функцию. Если использовать глобальную переменную при многопоточном программировании могут произойти конфликты
//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}
package main

import "fmt"

func createHugeString(end int) string {
	var s string
	for i := 0; i < end; i++ {
		s += "🥸"
	}
	return s
}

func someFunc() []rune {
	// создаем строку из 1024 символов
	v := createHugeString(1 << 10)
	//переводим строку в срез рун, чтобы при обрезании мы вернули нужное кол-во символов
	//если создать строку из символов которые занимают более 1 байта, то строка обрежется неверно
	res := []rune(v)
	//обрезаем строку
	return res[:100]
}

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
func main() {
	// Создаем локальную переменную вместо глобальной
	justString := someFunc()
	fmt.Println(string(justString))
}
