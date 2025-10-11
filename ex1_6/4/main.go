package main
// через runtime.Goexit()
import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("Defer executed")		
		for i := 0; i < 5; i++ {
			if i == 3 {
				fmt.Println("Calling runtime.Goexit()")
				runtime.Goexit()
			}
			fmt.Printf("Working... %d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("This won't execute")
	}()
	
	time.Sleep(3 * time.Second)
	fmt.Println("Main goroutine finished")
}