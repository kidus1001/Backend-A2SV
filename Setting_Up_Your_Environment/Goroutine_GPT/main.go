package main

import (
	"fmt"
	"time"
)

func printNums() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number ", i+1)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printNums()
	fmt.Println("Finished printing")

	time.Sleep(3 * time.Second)
}
