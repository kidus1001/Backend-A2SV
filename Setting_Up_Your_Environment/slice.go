package main

import (
	"fmt"
	"slices"
)

func slice() {
	var s []string
	fmt.Println("Unit:", s, s == nil, len(s) == 0)

	t := make([]int, 4, 5)
	fmt.Println(t, cap(t))

	t = append(t, 200)
	t = append(t, 200, 300, 400)
	fmt.Println(t)

	c := make([]int, 8, 50)
	copy(c, t)
	fmt.Println(c)

	if slices.Equal(t, c) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
