/*Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.

Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел. Для выбора опорного элемента можно взять середину или первый элемент.*/

package main

import "fmt"

func main() {
	arr := []int{12, 1993, 22, 1, -22, 4}
	sorted := Quicksort(arr)
	fmt.Println(sorted)
}

func Quicksort(input []int) []int {
	l := len(input)
	if l < 2 {
		return input
	}

	less := make([]int, 0)
	bigger := make([]int, 0)
	start := input[0]
	for _, v := range input[1:] {
		if v > start {
			bigger = append(bigger, v)
		} else {
			less = append(less, v)
		}
	}
	input = append(Quicksort(less), start)
	input = append(input, Quicksort(bigger)...)
	return input
}
