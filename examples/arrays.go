package examples

import "fmt"

func Arrays() {
	// Array will hold exactly 5 ints
	// both the type of items and length are part of array's type
	// empty arrays are initialized with 0's for ints
	var a [5]int
	fmt.Println("Empty array", a)

	// Set a value of an index
	fmt.Println("Old Value of a[2]", a[2])
	a[2] = 4

	// An error will occur if a different type is set as an array item, here its float64
	// a[4] = 2.1
	fmt.Println("New Value of a[2]", a[2])

	// builtin len function can be used to get the length of an array
	fmt.Println("Length of array a", len(a))

	// Array initializers
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array initialized during declaration", b)

	// 2d array initializer
	c := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("2d array", c)

	// Looping through a 2d array
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			fmt.Println(c[i][j])
		}
	}
}
