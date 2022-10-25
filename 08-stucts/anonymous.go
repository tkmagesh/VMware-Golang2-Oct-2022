package main

import "fmt"

func main() {
	s := struct {
		Id   int
		Name string
	}{100, "Magesh"}
	/*
		fmt.Println(s.Id)
		fmt.Println(s.Name)
	*/
	display(s)
}

func display(x struct {
	Id   int
	Name string
}) {
	fmt.Println(x.Id)
	fmt.Println(x.Name)
}
