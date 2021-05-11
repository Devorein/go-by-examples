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

// value receiver type wont reference the passed struct
func (r rect) perimeter() int {
	// Changing the fields of struct, but it wont have any effect since pointer is not used
	r.width = 100
	r.height = 100
	return 2*r.width + 2*r.height
}

func Methods() {
	r := rect{width: 20, height: 10}
	fmt.Println("area", r.area())
	fmt.Println("perimeter", r.perimeter())
	fmt.Println("new area", r.area())

	r2 := r
	fmt.Println("area", r2.area())
	fmt.Println("perimeter", r2.perimeter())
}
