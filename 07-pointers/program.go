package main

import "fmt"

func main() {
	//nos := [5]int{3, 1, 4, 2, 5} // => array
	/*
		nos := []int{3, 1, 4, 2, 5} // => slice
		newNos := nos
		fmt.Println(nos, newNos)
		nos[0] = 100
		fmt.Println(nos, newNos)
	*/

	var no int = 10
	//noPtr := &no //=> address of no
	/*
		var noPtr *int
		noPtr = &no
		fmt.Println(noPtr, no)

		//dereferencing
		newNo := *noPtr
		fmt.Println(newNo)
	*/

	fmt.Println("Before incrementing, no = ", no)
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	x, y := 100, 200
	fmt.Println("Before swapping : ", x, y)
	swap(&x, &y)
	fmt.Println("Before swapping : ", x, y)
}

func increment(x *int) {
	*x++
}

func swap(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}
