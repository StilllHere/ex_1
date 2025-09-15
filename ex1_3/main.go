/*Реализовать постоянную запись данных в канал (в главной горутине).

Реализовать набор из N воркеров, которые читают данные из этого канала и
выводят их в stdout.

Программа должна принимать параметром количество воркеров и
при старте создавать указанное число горутин-воркеров.*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N int = 7

func main() {
	channel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(N)

	for range make([]struct{}, N) {
		go Worker(&wg, channel)
	}
	Writer(channel)
	wg.Wait()
}

func Writer(ch chan<- int) {
	for {
		data := rand.Intn(100)
		ch <- data
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
