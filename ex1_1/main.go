/*
Реализовать встраивание методов

	в структуре Action от родительской структуры Human (аналог наследования).
*/
package main

import "fmt"

type Human struct {
	name    string
	surname string
	age     int
}

func (A Human) jogger() string {
	return fmt.Sprintf("%s %s who is jogger", A.name, A.surname)
}

type Action struct {
	Human
	duration         float64
	type_of_acrivity string
}

func main() {
	person1 := Action{
		Human: Human{
			"Ellie",
			"White",
			44,
		},
		duration:         1.5,
		type_of_acrivity: "running",
	}
	fmt.Println(person1.jogger())
}
