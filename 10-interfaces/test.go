package main

import "fmt"

type MyInterface interface {
	setData(int)
}

type MyStruct struct {
	data int
}

//MyInterface implementation
func (myStruct *MyStruct) setData(data int) {
	fmt.Printf("[setData] &myStruct = %p\n", myStruct)
	myStruct.data = data
}

func changeValue(i MyInterface, val int) {
	fmt.Printf("[changeValue] &i = %p\n", &i)
	i.setData(val)
}

func main() {
	s := MyStruct{0}
	fmt.Printf("[main] &myStruct = %p\n", &s)
	changeValue(&s, 100)
	fmt.Println(s.data)
}
