package main

import "fmt"

func sums(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func variadic() {
	fmt.Println(sums(1, 2, 3))
}
