package main

import (
	"fmt"
	"sync"
	"time"
)

//var wg sync.WaitGroup

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go f1(wg) //scheduling the function execution
	go f2(wg)
	wg.Wait() // counter to become 0
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrements the counter by 1
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
}

func f2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("f2 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f2 completed")
}
