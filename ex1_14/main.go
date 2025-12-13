/*Разработать программу, которая в runtime способна определить тип переменной, переданной в неё (на вход подаётся interface{}). Типы, которые нужно распознавать: int, string, bool, chan (канал).

Подсказка: оператор типа switch v.(type) поможет в решении.*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	//int
	reflectCheck(int8(1))
	reflectCheck(uint64(99999))
	//string
	reflectCheck("строка")
	//bool
	reflectCheck(true)
	reflectCheck(false)
	//channel
	channel := make(chan int)
	reflectCheck(channel)

	//int
	sprintCheck(int8(1))
	sprintCheck(uint64(99999))
	//string
	sprintCheck("строка")
	//bool
	sprintCheck(true)
	sprintCheck(false)
	//channel
	channel = make(chan int)
	sprintCheck(channel)
}

func reflectCheck(value interface{}) {
	fmt.Println(reflect.TypeOf(value))
}

func sprintCheck(value interface{}) {
	fmt.Println(fmt.Sprintf("%T", value))
}