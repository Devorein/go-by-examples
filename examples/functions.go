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

func Functions() {
	res := plus(1, 2, 3)
	fmt.Println(res)

	// Functions stored as value to variables
	fn := plus
	fmt.Println(fn(1, 2, 3))

	fmt.Println(fn2(1, "a"))
}
