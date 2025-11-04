package main

import (
	"fmt"
	"sync"
)

var count = 0
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	mu.Lock()
	count++
	mu.Unlock()
	//Something to tell the main function we are done
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Done! Count = ", count)
}
