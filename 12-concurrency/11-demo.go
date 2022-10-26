package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	result := <-ch // initiate the Receive operation
	fmt.Println(result)
}
