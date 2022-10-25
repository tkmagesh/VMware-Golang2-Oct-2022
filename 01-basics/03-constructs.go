package main

import "fmt"

func main() {
	/* if else */
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	if no := 21; no%2 == 0 { // no variable is not accessible outside the if else block
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	/* switch case */
	/*
		Rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Very Good"
		score 10 => "Excellent"
		otherwise => "Invalid Score"
	*/
	score := 6
	/*
		switch score {
		case 0:
			fmt.Println("Terrible")
		case 1:
			fmt.Println("Terrible")
		case 2:
			fmt.Println("Terrible")
		case 3:
			fmt.Println("Terrible")
		case 4:
			fmt.Println("Not Bad")
		case 5:
			fmt.Println("Not Bad")
		case 6:
			fmt.Println("Not Bad")
		case 7:
			fmt.Println("Not Bad")
		case 8:
			fmt.Println("Very Good")
		case 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid Score")
		}
	*/
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid Score")
	}

	/*
		no := 22
		switch {
		case no%2 == 0:
			fmt.Printf("%d is an even number\n", no)
		case no%2 != 0:
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	// mimicing if - else if - else if
	switch no := 21; {
	case no%2 == 0:
		fmt.Printf("%d is an even number\n", no)
	case no%2 != 0:
		fmt.Printf("%d is an odd number\n", no)
	}

	fmt.Println("Switch case with fallthrough")

	switch n := 2; n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
	case 7:
		fmt.Println("is <= 7")
		fallthrough
	case 8:
		fmt.Println("is <= 8")
	}

	//fallthrough usecase
	switch plan := "premium"; plan {
	case "super":
		fmt.Println("All super features")
		fallthrough
	case "premium":
		fmt.Println("All premium features")
		fallthrough
	case "pro":
		fmt.Println("All pro features")
		fallthrough
	case "free":
		fmt.Println("All free features")
	}

	/* for */
	fmt.Printf("\nfor\n")
	fmt.Println("for (v1.0)")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("for [while] (v2.0)")
	/*
		sum := 1
		for sum < 100 {
			sum += sum
		}
		fmt.Printf("sum = %d\n", sum)
	*/

	for sum := 1; sum < 100; {
		sum += sum
		fmt.Printf("sum = %d\n", sum)
	}

	fmt.Println("for [infinite] (v3.0)")
	numSum := 1
	for {
		numSum += numSum
		if numSum > 100 {
			break
		}
	}
	fmt.Printf("numSum = %d\n", numSum)

	fmt.Println("Labels")

OUTER_LOOP:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				fmt.Println("=================")
				continue OUTER_LOOP
			}
		}
	}

}
