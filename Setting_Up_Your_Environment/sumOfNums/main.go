package main

import (
	"fmt"
)

func controller(nums *[]int) {
	fmt.Println("Enter the size for your array: ")
	var size int
	fmt.Scanln(&size)

	for i := 0; i < size; i++ {
		fmt.Print("Enter a number: ")
		var input int
		fmt.Scanln(&input)
		*nums = append(*nums, input)
	}
}

func calculator(nums *[]int) int {
	if len(*nums) == 0 {
		return 0
	}
	total := 0
	for _, i := range *nums {
		total += i
	}
	return total
}

func main() {
	var nums []int
	controller(&nums)
	var total = calculator(&nums)
	fmt.Println("Sum of numbers = ", total)
}
