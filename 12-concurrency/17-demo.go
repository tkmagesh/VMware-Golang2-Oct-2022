package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	close(ch)
	go func() {
		fmt.Println("Send Attempted")
		ch <- 100
		fmt.Println("Send completed")
	}()
	time.Sleep(4 * time.Second)
	fmt.Println("Done")
}
