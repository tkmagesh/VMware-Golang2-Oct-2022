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
	ch := genPrimes(3)
	for primeNo := range ch {
		fmt.Println(primeNo)
	}
	fmt.Println("Done")
}

func genPrimes(start int) <-chan int {
	ch := make(chan int)

	timeoutCh := func() <-chan time.Time {
		timeoutCh := make(chan time.Time)
		go func() {
			time.Sleep(10 * time.Second)
			timeoutCh <- time.Now()
		}()
		return timeoutCh
	}()

	go func() {
		no := start
		for {
			<-timeoutCh
			if isPrime(no) {
				time.Sleep(500 * time.Millisecond)
				ch <- no
			}
			no++
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
