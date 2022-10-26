/*
Write a goroutine that generates prime number within the given start & end
the main function will print the generated prime numbers
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})
	ch := genPrimes(3, stop)

	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		stop <- struct{}{}
	}()

	for primeNo := range ch {
		fmt.Println(primeNo)
	}
	fmt.Println("Done")
}

func genPrimes(start int, stop chan struct{}) <-chan int {
	ch := make(chan int)

	go func() {
		no := start
	LOOP:
		for {
			select {
			case <-stop:
				break LOOP
			default:
				if isPrime(no) {
					time.Sleep(500 * time.Millisecond)
					ch <- no
				}
				no++
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
