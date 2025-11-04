package main

import "fmt"

type person struct {
	name  string
	age   int
	hired bool
}

func newPerson(name string, age int) *person {
	p := person{name: name, age: age}
	p.hired = false
	return &p
}

func structs() {
	p := person{"Kidus", 23, true}
	fmt.Println(p)

	person1 := person{name: "Kidus", age: 23}
	fmt.Println(person1)

	person2 := newPerson("KidusYosef", 26)
	fmt.Println(person2)

	dog := struct {
		name       string
		isgood     bool
		discipline string
	}{
		"Jackie",
		true,
		"A+",
	}

	fmt.Println(dog)
}
