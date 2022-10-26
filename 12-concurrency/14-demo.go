package main

import "fmt"

func main() {
	ch := make(chan int)
	go fn(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fn(ch chan int) {
	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40
	ch <- 50
}

/*
func main() {
	ch := make(chan int)
	count := 10
	go fn(ch, count)
	for i := 0; i < (count); i++ {
		fmt.Println("	Receive attempted ", (i+1)*10)
		fmt.Println(<-ch)
		fmt.Println("	Receive succeeded ", (i+1)*10)
	}
}

func fn(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Println("	Send attempted ", (i)*10)
		ch <- i * 10
		fmt.Println("	Send succeeded ", (i)*10)
	}
}
*/
