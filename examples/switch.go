package examples

import (
	"fmt"
	"time"
)

func Switch() {
	i := 2
	fmt.Print("Write", i, "as")

	// cases don't need to be enclosed
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Multiple expressions in a single case, separated by comma
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// switch statements without an expression
	// non constants expression as case condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before the noon")
	default:
		fmt.Println("It's after noon")
	}

	whichType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a boolean")
		case int:
			fmt.Println("I'm an integer")
		case float32:
			fmt.Println("I'm a floting point value")
		default:
			fmt.Printf("Unknown type %T\n", t)
		}
	}

	whichType(true)
	whichType(1)
	whichType(float32(1.2))
}
