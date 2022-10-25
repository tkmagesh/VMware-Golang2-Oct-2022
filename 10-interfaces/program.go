package main

import (
	"fmt"
	"math"
)

//3rd party libary
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

//your application
type ShapeWithArea interface {
	Area() float32
}

func PrintArea(x ShapeWithArea) {
	fmt.Printf("Area = %v\n", x.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func PrintPerimeter(x ShapeWithPerimeter) {
	fmt.Printf("Perimeter = %v\n", x.Perimeter())
}

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func PrintShape(s Shape) {
	PrintArea(s)
	PrintPerimeter(s)
}

func main() {
	c := Circle{12}
	//fmt.Printf("Area = %v\n", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{10, 12}
	//fmt.Printf("Area = %v\n", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
}
