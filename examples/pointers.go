package examples

import "fmt"

func zeroVal(val int) {
	val = 0
}

func zeroPtr(val *int) {
	*val = 0
}

func Pointers() {
	i := 1
	fmt.Println("Before", i)

	zeroVal(i)
	fmt.Println("zeroVal", i)

	zeroPtr(&i)
	fmt.Println("zeroPtr", i)

	fmt.Println("Pointer", &i)
}
