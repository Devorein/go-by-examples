package examples

import "fmt"

func Variables() {

	// Go infers types during initialization
	var a = "initial"
	fmt.Println(a)

	var d = true
	fmt.Println(d)

	// Multiple variable declaration in a single statement
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Uninitialized declarations are initialized with zero values, 0 for int
	var e int
	fmt.Println(e)

	// := declaring and initializing variables, var f string = "apple"
	f := "apple"
	fmt.Println(f)
}
