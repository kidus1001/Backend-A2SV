package main

import (
	"fmt"
	"math"
)

func constants() {
	const name = "Kidus"
	fmt.Println(name)

	const num = 50000
	fmt.Println(num)

	const bigNum float64 = 3e20
	fmt.Println(bigNum)

	var sqrRoot = math.Sqrt(25.0)
	fmt.Println(sqrRoot)
}
