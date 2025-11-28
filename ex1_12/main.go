/*
Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.
*/
package main

import "fmt"

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _,v := range slice {
		set[v] = struct{}{}
	}
	var result []string
	for key := range set {
		result = append(result, key)
	}
	fmt.Println(result)
}