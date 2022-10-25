package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	Product
	//Dummy
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{Id: id, Name: name, Cost: cost},
		Expiry:  expiry,
	}
}

func main() {
	//apple := PerishableProduct{Product: Product{Id: 102, Name: "Apple", Cost: 12}, Expiry: "5 Days"}
	apple := NewPerishableProduct(102, "Apple", 12, "5 Days")
	fmt.Println(apple)
	fmt.Println(apple.Product.Name)
	fmt.Println(apple.Name)
}
