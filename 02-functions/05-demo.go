/* Higher order functions - functions as return values */

package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		fn := getFn()
		fn()
	*/

	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	/*
		profiledAdd := getProfiledOperation(add)
		profiledAdd(100, 200)

		profiledSubtract := getProfiledOperation(subtract)
		profiledSubtract(100, 200)
	*/

	/*
		logAdd := getLogOperation(add)
		profiledLogAdd := getProfiledOperation(logAdd)
	*/
	profiledLogAdd := getProfiledOperation(getLogOperation(add))
	profiledLogAdd(100, 200)

	profiledLogSubtract := getProfiledOperation(getLogOperation(subtract))
	profiledLogSubtract(100, 200)

}

func getProfiledOperation(op func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		op(x, y)
		elapsed := time.Now().Sub(start)
		fmt.Printf("Time taken : %v\n", elapsed)
	}
}
func getLogOperation(op func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("op commenced")
		op(x, y)
		fmt.Println("op completed")
	}
}

//3rd party library
func add(x, y int) {
	fmt.Println("Add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}

func getFn() func() {
	return func() {
		fmt.Println("anon fn invoked")
	}
}
