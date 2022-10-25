package main

import "fmt"

func main() {
	nos := make([]int, 0, 10)
	//nos := []int{10}
	nos = append(nos, 10)
	fmt.Printf("len = %d, capacity = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 20)
	fmt.Printf("len = %d, capacity = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 30)
	fmt.Printf("len = %d, capacity = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 40)
	fmt.Printf("len = %d, capacity = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 50)
	fmt.Printf("len = %d, capacity = %d, nos = %v\n", len(nos), cap(nos), nos)
}
