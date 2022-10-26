package main

import (
	"fmt"
	calc "modularity-demo/calculator"
	"modularity-demo/calculator/utils"

	"github.com/fatih/color"
	/* _ "modularity-demo/calculator" */)

func main() {
	color.Green("app executed")
	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println(calculator.OpCount())
	*/

	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println(calc.OpCount())
	fmt.Println(utils.IsEven(21))
	calc.ProductRanks["Pen"] = 1
	fmt.Println(calc.ProductRanks)

}
