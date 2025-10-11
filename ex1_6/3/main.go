package main
// через уведомления
import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	end := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1) 
	
	go func() {
		defer wg.Done()
		for {
			select {
			case q := <-ch:
				fmt.Println(q)
			case <-end:
				fmt.Println("goroutine is done")
				return
			}
		}
	}()

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	end <- true
	
	wg.Wait() 
	fmt.Println("the end")
}