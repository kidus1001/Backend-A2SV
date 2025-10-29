package main

func closure() func() int {
	number := 0

	return func() int {
		number += 1
		return number
	}
}
