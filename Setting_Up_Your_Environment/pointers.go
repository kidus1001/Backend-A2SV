package main

func byValue(num int) {
	num = 0
}

func pointers(num *int) {
	*num = 0
}
