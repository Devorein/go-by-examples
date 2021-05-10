package examples

import "fmt"

func intSequence(startValue int) func(increment int) int {
	i := startValue

	return func(increment int) int {
		i += increment
		return i
	}
}

func Closures() {
	// Creating a closure
	nextInt := intSequence(0)
	fmt.Println(nextInt(1))
	fmt.Println(nextInt(2))
	fmt.Println(nextInt(3))
}
