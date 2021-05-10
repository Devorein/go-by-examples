package examples

import "fmt"

func Values() {

	// String concatenation
	fmt.Println("go" + "lang")

	// Println takes in a variadic number of arguments, each separated by a single space
	fmt.Println("1+1 =", 1+1)

	// Integer division
	fmt.Println("7/3 =", 7/3)

	// Float division since one of the operand is a float32
	fmt.Println("7.0/3 =", 7.0/3)

	// Float division
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Boolean expressions
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
