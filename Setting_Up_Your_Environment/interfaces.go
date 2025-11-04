package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float32
	perimeter() float32
}

type rect struct {
	width, height float32
}

type circle struct {
	radius float32
}

func (r rect) area() float32 {
	return r.height * r.width
}

func (r rect) perimeter() float32 {
	return 2*r.height + 2*r.width
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func deleteCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("Deleted circle with radius of ", c.radius)
	}
}

func interfaces() {
	c := circle{2.1}
	r := rect{20, 2}

	measure(c)
	measure(r)

	deleteCircle(c)
	deleteCircle(r)
}
