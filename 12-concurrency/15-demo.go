package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	for {
		time.Sleep(500 * time.Millisecond)
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
}

func fn(ch chan int) {
	count := 10
	for i := 1; i <= count; i++ {
		ch <- i * 10
	}
	close(ch)
}
