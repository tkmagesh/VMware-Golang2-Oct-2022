/* sharing state */
/*
	to detect data race
		go run --race <program.go>
		go build --race <program.go>
*/
package main

import (
	"fmt"
	"sync"
)

var counter = 0

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn(wg)
	}
	wg.Wait()
	fmt.Println("Counter =", counter)
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	counter++
}
