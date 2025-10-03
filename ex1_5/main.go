package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	duration := 3 * time.Second

	go func() {
		for value := range ch {
			fmt.Printf("Received: %d\n", value)
		}
		fmt.Println("Reader stopped")
	}()

	go func() {
		timer := time.After(duration)
		counter := 0

		for {
			counter++
			select {
			case ch <- counter:
				time.Sleep(100 * time.Millisecond)
			case <-timer:
				close(ch)
				fmt.Println("Writer stopped")
				return
			}
		}
	}()

	time.Sleep(duration + 100*time.Millisecond)
	fmt.Println("Program finished")
}
