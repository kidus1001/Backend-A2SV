package main

import "fmt"

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, " ", c)

	var d = true
	fmt.Println(!d)

	e := "Mama"
	fmt.Println(e)

	var f int
	fmt.Println(f)

	var g string
	fmt.Println(g)
}
