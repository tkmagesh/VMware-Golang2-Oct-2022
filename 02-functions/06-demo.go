package main

import "fmt"

func main() {
	operations := []func(int, int){
		func(i1, i2 int) {
			fmt.Println("Add result : ", i1+i2)
		},
		func(i1, i2 int) {
			fmt.Println("Subtract result : ", i1-i2)
		},
		func(i1, i2 int) {
			fmt.Println("Multiply result : ", i1*i2)
		},
		func(i1, i2 int) {
			fmt.Println("Divide result : ", i1/i2)
		},
	}

	for _, op := range operations {
		op(100, 200)
	}
}
