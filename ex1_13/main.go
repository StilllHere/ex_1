
/*
Поменять местами два числа без использования временной переменной.

Подсказка: примените сложение/вычитание или XOR-обмен.
*/
package main

import "fmt"

func main() {
	a, b := 22, 43
	fmt.Printf("a:%d, b:%d\n", a, b)

	a, b = b, a
	fmt.Printf("a:%d, b:%d", a, b)
}