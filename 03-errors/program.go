package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be 0")

func main() {
	/*
		quotient, remainder, err := divide(100, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
	*/

	quotient, remainder, err := divide(100, 0)
	if errors.Is(err, DivideByZeroError) {
		fmt.Println("do not attempt to divide by zero")
		return
	}
	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be 0")
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

/*
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("divisor cannot be 0")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
