/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2. После этого данные из второго канала должны выводиться в stdout. То есть,
 организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка. 
Убедитесь, что чтение из второго канала корректно завершается.*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		defer close(ch1)

		for i := range arr {
			ch1 <- arr[i]
		}

	}()

	go func() {
		defer wg.Done()
		defer close(ch2)
		for q := range ch1 {
			q *= 2
			ch2 <- q
		}
	}()

	go func() {
		defer wg.Done()
		for result := range ch2 {
			fmt.Println(result)
		}
	}()
	wg.Wait()
}