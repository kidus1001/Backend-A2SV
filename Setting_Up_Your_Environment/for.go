package main

import "fmt"

func forFunc() {
	var i int = 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("Range : ", i)
	}

	for {
		fmt.Println("Loop")
		break
	}

	for x := range 6 {
		if x%2 == 0 {
			fmt.Println(x)
		} else {
			continue
		}
	}
}
