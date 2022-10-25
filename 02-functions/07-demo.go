/* variadic functions */

package main

import "fmt"

func main() {
	fmt.Println(sum(10))             //=> 10
	fmt.Println(sum(10, 20))         //=> 30
	fmt.Println(sum(10, 20, 30, 40)) //=> 100
}

func sum(nos ...int) int {
	result := 0
	for i := 0; i < len(nos); i++ {
		result += nos[i]
	}
	return result
}
