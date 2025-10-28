package main

import (
	"fmt"
)

func arrays() {
	var a [5]int
	fmt.Println("emp: ", a)

	a[1] = 100
	fmt.Println(a)

	fmt.Println(a[1], len(a))

	b := [10]int{1, 2, 3}

	fmt.Println(b)

	c := [...]int{1, 2, 3}
	fmt.Println(c)

	d := [...]int{100, 7: 400, 500}
	fmt.Println(d)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
