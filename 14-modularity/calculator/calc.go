package calculator

import "fmt"

var operation_count int

var ProductRanks map[string]int

func init() {
	ProductRanks = make(map[string]int)
	fmt.Println("calculator-init[1] triggered")
}

func init() {
	fmt.Println("calculator-init[2] triggered")
}

func add_operation_count() {
	operation_count++
}

func OpCount() int {
	return operation_count
}
