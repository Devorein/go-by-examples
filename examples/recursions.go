package examples

import "fmt"

func factorial(a int) int {
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func Recursions() {
	fmt.Println(factorial(4))
	fmt.Println(fibonacci(4))
}
