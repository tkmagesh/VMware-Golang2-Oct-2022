package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 10
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 20
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println(<-ch3)
	}()

	for i := 0; i < 3; i++ {
		select {
		case data1 := <-ch1:
			fmt.Println("Ch-1 data : ", data1)
		case data2 := <-ch2:
			fmt.Println("Ch-2 data : ", data2)
		case ch3 <- 100:
			fmt.Println("Data sent to Ch3")
			/*
				default:
					fmt.Println("Default case triggered")
			*/
		}
	}

}
