package main

func recursion(n int) int {
	if n == 1 {
		return 1
	}
	return n * recursion(n-1)
}
