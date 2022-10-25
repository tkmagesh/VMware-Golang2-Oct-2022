package main

import "fmt"

func main() {
	//var x interface{}
	var x any
	x = 100
	x = "This is a string"
	x = true
	x = struct{}{}
	x = 100.765
	fmt.Println(x)

	var no interface{}
	no = "100"
	//y := no + 200

	//NOT SAFE
	//y := no.(int) + 200

	if val, ok := no.(int); ok {
		result := val + 200
		fmt.Println(result)
	} else {
		fmt.Println("no is not an int")
	}
	//fmt.Println(y)

	x = 100.23
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 = ", val+200)
	case string:
		fmt.Println("x is a string, len(x) = ", len(val))
	case bool:
		fmt.Println("x is bool, !x = ", !val)
	default:
		fmt.Println("unknown type")
	}
}
