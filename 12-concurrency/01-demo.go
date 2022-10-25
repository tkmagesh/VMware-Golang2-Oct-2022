package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling the function execution
	f2()
	//fmt.Scanln()
	time.Sleep(1 * time.Second)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
