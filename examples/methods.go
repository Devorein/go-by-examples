package examples

import "fmt"

type rect struct {
	width  int
	height int
}

// Methods of struct rect
func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func Methods() {
	r := rect{width: 20, height: 10}
	fmt.Println("area", r.area())
	fmt.Println("perimeter", r.perimeter())

	r2 := r
	fmt.Println("area", r2.area())
	fmt.Println("perimeter", r2.perimeter())
}
