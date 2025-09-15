package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

const N int = 7

func main() {
	channel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(N)
	notifyContext, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	
	for range make([]struct{}, N) {
		go Worker(&wg, channel)
	}
	Writer(notifyContext, channel)
	wg.Wait()
}

func Writer(ctx context.Context, ch chan<- int) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			fmt.Println("user did ctrl+c")
			return
		default:
			data := rand.Intn(10000)
			ch <- data
		}
	}
}

func Worker(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	for data := range c {
		fmt.Println(data)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println("worker is finished")
}