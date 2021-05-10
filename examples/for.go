package examples

import "fmt"

func For() {

	// initialization done outside for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		// Updating the loop variables done inside for loop
		i = i + 1
	}

	// Classic initialization/condition/update pattern of for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// An infinite loop, only broken by a break/return statement
	for {
		fmt.Println("loop")
		break
	}

	// continue statement to skip the rest of the for loop body for current iteration and move to the next iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
