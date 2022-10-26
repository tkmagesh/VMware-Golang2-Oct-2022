package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go fn(ch, wg)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Closing the channel")
		close(ch)
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	wg.Wait()
}

func fn(ch chan int, wg *sync.WaitGroup) {
	count := 10
	for i := 1; i <= count; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Send attempted : ", i*10)
		ch <- i * 10
		fmt.Println("Send successful : ", i*10)
	}
	close(ch)
	wg.Done()
}
