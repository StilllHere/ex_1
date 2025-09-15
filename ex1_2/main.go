package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _,n := range numbers {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n * n)
			wg.Done()
		}(n)
	}
	wg.Wait()
}