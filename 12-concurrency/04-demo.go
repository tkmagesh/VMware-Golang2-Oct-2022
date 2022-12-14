package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("Hit ENTER to start...")
	fmt.Scanln()
	var goRoutineCount = flag.Int("count", 1, "Number of goroutines to run")
	flag.Parse()
	rand.Seed(7)
	for i := 1; i <= *goRoutineCount; i++ {
		wg.Add(1)
		go fn(i, time.Duration(rand.Intn(10))*time.Second, wg)
	}
	wg.Wait()
	fmt.Println("Done")
}

func fn(id int, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("fn [ id = %d ] started\n", id)
	time.Sleep(delay)
	fmt.Printf("fn [ id = %d ] completed\n", id)
}
