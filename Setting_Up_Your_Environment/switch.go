package main

import (
	"fmt"
	"reflect"
	"time"
)

func switchFunc() {
	num := 3
	fmt.Print("Write ", num, " as ")
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int", t)
		default:
			fmt.Println("Unsure about my type!", reflect.TypeOf(i))
		}
	}
	whatAmI(20.2)
}
