package calculator

import "fmt"

func init() {
	fmt.Println("calculator/subtract.go init triggered")
}

func Subtract(x, y int) int {
	//operation_count++
	add_operation_count()
	return x - y
}
