package calculator

import "fmt"

func init() {
	fmt.Println("calculator/add.go init triggered")
}

func Add(x, y int) int {
	//operation_count++
	add_operation_count()
	return x + y
}
