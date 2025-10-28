package main

import (
	"fmt"
)

func mapsFunc() {
	m := make(map[string]int)

	m["Cristiano"] = 7
	m["Messi"] = 10

	fmt.Println("Mao:", m)
	fmt.Println(m["Ma"])

	delete(m, "Cristiano")

	fmt.Println(m)
}
