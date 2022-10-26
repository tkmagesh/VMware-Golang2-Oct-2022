package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	f1Ctx, cancel := context.WithTimeout(rootCtx, 10*time.Second)
	defer cancel() //to avoid any memory leak

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go f1(f1Ctx, wg)
	wg.Wait()
}

func f1(ctx context.Context, wg *sync.WaitGroup) {

	fmt.Println("f1 started")

	fnWg := &sync.WaitGroup{}
	for i := 1; i <= 2; i++ {
		fnWg.Add(1)
		go fn(i, ctx, fnWg)
	}

LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("f1 - cancel signal received")
			break LOOP
		default:
			time.Sleep(300 * time.Millisecond)
			fmt.Println("f1 - doing something")
		}
	}
	fmt.Println("f1 completed")
	fnWg.Wait()
	wg.Done()
	return
}

func fn(id int, ctx context.Context, wg *sync.WaitGroup) {
	fmt.Printf("fn[%d] started\n", id)

LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("fn[%d] - cancel signal received\n", id)
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("fn[%d] - doing something\n", id)
		}
	}
	fmt.Printf("fn[%d] completed\n", id)
	wg.Done()

}
