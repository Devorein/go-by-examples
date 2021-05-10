package examples

import "fmt"

func Range() {
	var sum int = 0
	arr := [5]int{1, 2, 3, 4, 5}

	// Range in context of an array, index, value tuple
	for i, v := range arr {
		fmt.Println("Index", i, "Value", v)
		sum += v
	}

	fmt.Println(sum)

	// Range in context of a map, key, value tuple
	maps := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for k, v := range maps {
		fmt.Println("Key", k, "Value", v)
	}

}
