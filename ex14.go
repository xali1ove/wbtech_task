//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.
package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Рефлексия позволяет получить информацию о типе переменной во время выполнения программы.
	var v1 interface{} = 42
	var v2 interface{} = "hello"
	var v3 interface{} = true
	var v4 interface{} = make(chan int)

	fmt.Println("Type of v1:", reflect.TypeOf(v1))
	fmt.Println("Type of v2:", reflect.TypeOf(v2))
	fmt.Println("Type of v3:", reflect.TypeOf(v3))
	fmt.Println("Type of v4:", reflect.TypeOf(v4))
}

/*Основные функции и методы пакета reflect:

reflect.TypeOf(v interface{}): возвращает тип переменной v.
reflect.ValueOf(v interface{}): возвращает значение переменной v.
reflect.New(t Type): создает новый экземпляр переменной заданного типа t.
reflect.ValueOf(v interface{}).Elem(): возвращает указатель на переменную v.
Value.Kind(): возвращает базовый тип переменной, например, int, string, bool, slice, map, struct и т.д.
Value.Interface(): возвращает значение переменной в виде пустого интерфейса interface{}.
Value.Field(i int): возвращает значение поля структуры с индексом i.
Value.MethodByName(name string): возвращает метод с именем name.
*/
