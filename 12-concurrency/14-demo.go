package main

import "fmt"

func main() {
	ch := make(chan int)
	go fn(ch)
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

}

func fn(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i * 10
	}
}
