package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func main() {
	//var pen Product
	//var pen Product = Product{100, "Pen", 5}
	//var pen = Product{100, "Pen", 5}
	//pen := Product{100, "Pen", 5}
	pen := Product{Id: 100, Name: "Pen"}
	fmt.Printf("%#v\n", pen)

	/*
		p1 := pen
		p1.Cost = 5
		fmt.Printf("%#v\n", p1)
	*/
	p1 := Product{Id: 100, Name: "Pen"}
	fmt.Println(pen == p1)
}
