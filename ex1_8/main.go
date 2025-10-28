/*Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.

Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).

Подсказка: используйте битовые операции (|, &^).*/

package main

import "fmt"

func main() {
	number:=int64(5)
	result1:=setbit(number,0)
	result2:=clearbit(number,0)
	fmt.Println(result1)
	fmt.Println(result2)
}

func setbit(n int64, pos uint) int64 {
	n |= int64(1 << pos)
	return n
}

func clearbit(n int64, pos uint) int64 {
	mask := ^int64(1 << pos)
	n &= mask
	return n
}
