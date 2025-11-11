package main

import "fmt"

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	rez:= make(map[int][]float64)

	for _, v := range data {
		fullPart := int(v/10) * 10
		rez[fullPart] = append(rez[fullPart], v)
	}

	for key, value := range rez{
		fmt.Println(key, value)
	}
}