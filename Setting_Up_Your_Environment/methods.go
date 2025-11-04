package main

import "fmt"

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.height * r.width
}

func (r rectangle) perimeter() int {
	return (2 * r.height) + (2 * r.width)
}

func methods() {
	r := rectangle{10, 5}
	area := r.area()
	perim := r.perimeter()

	fmt.Println(area)
	fmt.Println(perim)

	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perimeter())

}
