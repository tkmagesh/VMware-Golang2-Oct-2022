package main

import (
	"fmt"
)

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (Product) WhoAmI() {
	fmt.Println("I am a product")
}

func (product Product) Print() {
	fmt.Printf("Id = %d, Name = %q, Cost = %v\n", product.Id, product.Name, product.Cost)
}

func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v", product.Id, product.Name, product.Cost)
}

func Display(product Product) {
	fmt.Printf("Id = %d, Name = %q, Cost = %v\n", product.Id, product.Name, product.Cost)
}

/*
func ApplyDiscount(product *Product, percentage float32) {
	product.Cost = product.Cost * ((100 - percentage) / 100)
}
*/

func (product *Product) ApplyDiscount(percentage float32) {
	product.Cost = product.Cost * ((100 - percentage) / 100)
}

type PerishableProduct struct {
	Product
	//Dummy
	Expiry string
}

//overriding the "Product.Print" method
func (pp PerishableProduct) Print() {
	fmt.Printf("Id = %d, Name = %q, Cost = %v, Expiry = %q\n", pp.Id, pp.Name, pp.Cost, pp.Expiry)
}

func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%v, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{Id: id, Name: name, Cost: cost},
		Expiry:  expiry,
	}
}

func main() {
	pen := Product{100, "Pen", 5}
	//Print(pen)
	pen.Print()
	//pen.WhoAmI()

	fmt.Println("After applying 10% discount")
	//ApplyDiscount(&pen, 10)
	//(&pen).ApplyDiscount(10)
	pen.ApplyDiscount(10)
	pen.Print()

	apple := NewPerishableProduct(100, "apple", 10, "5 Days")

	//using a function
	Display(apple.Product)

	//using a method
	//apple.Print()
	fmt.Println(apple.Format())

	apple.ApplyDiscount(50)
	//apple.Print()
	fmt.Println(apple.Format())
}
