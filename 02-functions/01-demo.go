/* function basics */
package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	fmt.Println(getGreetMsg("John"))
	fmt.Printf("add(100,200) = %d\n", add(100, 200))

	//fmt.Println(divide(100, 7))
	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
	*/

	quotient, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d \n", quotient)

}

/* basic function */
func sayHi() {
	fmt.Println("Hi")
}

/* function with 1 argument */
func greet(userName string) {
	fmt.Printf("Hi %s, Have a good day!\n", userName)
}

/* function with 1 argument and 1 return result */
func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a good day!\n", userName)
}

/* function with 2 arguments and 1 return result */
/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

/* function with 2 arguments and 2 return results */
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
