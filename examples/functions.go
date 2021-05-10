package examples

import (
	"fmt"
	"strconv"
)

// Function that takes in a variadic number of integer arguments and returns an integer
func plus(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// takes in two arguments one int another string and returns a string
func fn2(a int, b string) string {
	return strconv.Itoa(a) + b
}

// Function returning tuples
func fn3() (int, int) {
	return 3, 6
}

func Functions() {
	res1 := plus(1, 2, 3)
	fmt.Println(res1)

	res2 := plus([]int{1, 3, 4}...)
	fmt.Println(res2)

	// Functions stored as value to variables
	fn := plus
	fmt.Println(fn(1, 2, 3))

	fmt.Println(fn2(1, "a"))

	// Initializing variables using function that returns a tuple
	a, b := fn3()
	fmt.Println(a, b)

	fmt.Println(fn3())

	// only extracting required values from a function that returns a tuple
	_, c := fn3()
	fmt.Println(c)
}
