package main

import "fmt"

func main() {
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println("Before sorting, nos : ", nos)
	sort(nos)
	fmt.Println("After sorting, nos : ", nos)
}

func sort(nos [5]int) {
	for i := 0; i < (len(nos) - 1); i++ {
		for j := i + 1; j < len(nos); j++ {
			if nos[i] > nos[j] {
				nos[i], nos[j] = nos[j], nos[i]
			}
		}
	}
}
