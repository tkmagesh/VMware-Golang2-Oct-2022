package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	f1Ctx, cancel := context.WithCancel(rootCtx)
	defer cancel() //to avoid any memory leak
	go func() {
		fmt.Println("Hit ENTER to stop..")
		fmt.Scanln()
		cancel()
	}()

	done := f1(f1Ctx)
	<-done
}

func f1(ctx context.Context) <-chan struct{} {
	done := make(chan struct{})
	fmt.Println("f1 started")
	go func() {
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
		done <- struct{}{}
	}()
	return done
}
