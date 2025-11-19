/*Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы, присутствующие и в первом, и во втором.

Пример:
A = {1,2,3}
B = {2,3,4}
Пересечение = {2,3} */

package main

import "fmt"

func main() {
	ar1 := []int{1, 2, 22}
	ar2 := []int{1, 2, 33}

	fmt.Println(intersect(&ar1, &ar2))

}

func intersect(ar1, ar2 *[]int) []int {
	var result []int
	set := make(map[int]struct{})
	for _, v := range *ar1 {
		set[v] = struct{}{}
	}
	for _, v := range *ar2 {
		_, yep := set[v]
		if yep {
			result = append(result, v)
		}
	}
	return result
}