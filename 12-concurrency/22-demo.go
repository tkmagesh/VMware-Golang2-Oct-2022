package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2) // buffered channel
	fmt.Println("len(ch) = ", len(ch))
	ch <- 100
	ch <- 200
	fmt.Println("len(ch) = ", len(ch))
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
