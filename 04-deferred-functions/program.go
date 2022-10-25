package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("	deferred @ main")
		if e := recover(); e != nil {
			fmt.Println("Panic encountered : ", e)
			return
		}
		fmt.Println("Thank you!")
	}()
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	defer func() {
		fmt.Println("	deferred @ f1 [1]")
	}()
	defer func() {
		fmt.Println("	deferred @ f1 [2]")
	}()
	defer func() {
		fmt.Println("	deferred @ f1 [3]")
	}()
	fmt.Println("f1 started")
	f2()
	fmt.Println("f1 completed")
}

func f2() {
	defer fmt.Println("	deferred @ f2 [1]")
	defer fmt.Println("	deferred @ f2 [2]")
	defer fmt.Println("	deferred @ f2 [3]")
	fmt.Println("f2 started")
	divisor := 7
	result := 100 / divisor
	fmt.Println(result)
	fmt.Println("f2 completed")
}
