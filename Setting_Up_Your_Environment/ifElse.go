package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ifElse() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(201)
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
