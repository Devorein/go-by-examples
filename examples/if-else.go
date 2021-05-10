package examples

import "fmt"

func IfElse() {

	// Basic if else statement
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Standalone if statement
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Declaring and initializing variable in if/else block, which is scoped to it
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
