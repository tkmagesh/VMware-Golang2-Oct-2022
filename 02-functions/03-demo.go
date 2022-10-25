/* higher order functions - functions can be assigned as values to variables */

package main

import "fmt"

type operation func(int, int) int

func main() {
	var fn func()
	fn = func() {
		fmt.Println("anonymous function invoked")
	}
	fn()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a good day!\n", userName)
	}
	greet("Magesh")

	var add func(int, int)
	add = func(x, y int) {
		fmt.Println(x + y)
	}
	add(100, 200)

	/*
		var subtract func(int, int) int
		subtract = func(x, y int) int {
			return x - y
		}
		subtractResult := subtract(100, 200)
		fmt.Println("subtract result =", subtractResult)

		var multiply func(int, int) int
		multiply = func(x, y int) int {
			return x * y
		}
		multiplyResult := multiply(100, 200)
		fmt.Println("multiply result =", multiplyResult)
	*/

	/*
		var subtract operation
		subtract = func(x, y int) int {
			return x - y
		}
		subtractResult := subtract(100, 200)
		fmt.Println("subtract result =", subtractResult)

		var multiply operation
		multiply = func(x, y int) int {
			return x * y
		}
		multiplyResult := multiply(100, 200)
		fmt.Println("multiply result =", multiplyResult)
	*/

	//var op func(int, int) int
	var op operation
	op = func(x, y int) int {
		return x - y
	}
	subtractResult := op(100, 200)
	fmt.Println("subtract result =", subtractResult)

	op = func(x, y int) int {
		return x * y
	}
	multiplyResult := op(100, 200)
	fmt.Println("multiply result =", multiplyResult)
}
