/* higher order functions - functions as arguments to other functions */

package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		exec(fn)
		exec(func() {
			fmt.Println("anon fn invoked")
		})
	*/

	/*
		fmt.Println("operation started")
		add(100, 200)
		fmt.Println("operation completed")

		fmt.Println("operation started")
		subtract(100, 200)
		fmt.Println("operation completed")
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(100, 200, add)
	logOperation(100, 200, subtract)

	profileOperation(100, 200, add)
	profileOperation(100, 200, subtract)
}

func logOperation(x, y int, op func(int, int)) {
	fmt.Println("op commenced")
	op(x, y)
	fmt.Println("op completed")
}

func profileOperation(x, y int, op func(int, int)) {
	start := time.Now()
	op(x, y)
	elapsed := time.Now().Sub(start)
	fmt.Printf("Time taken : %v\n", elapsed)
}

func logAdd(x, y int) {
	fmt.Println("operation commenced")
	add(x, y)
	fmt.Println("operation completed")
}

func logSubtract(x, y int) {
	fmt.Println("operation commenced")
	subtract(x, y)
	fmt.Println("operation completed")
}

//3rd party library
func add(x, y int) {
	fmt.Println("Add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}

func fn() {
	fmt.Println("fn invoked")
}

func exec(f func()) {
	f()
}
