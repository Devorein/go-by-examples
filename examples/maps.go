package examples

import "fmt"

func Maps() {
	m := make(map[string]int)

	m["1"] = 1
	m["2"] = 2

	fmt.Println(m)

	// Get a single value from a map using its key
	fmt.Println(m["2"])

	// Getting the length of a map
	fmt.Println(len(m))

	// delete a key/value pair from a map
	delete(m, "1")
	fmt.Println(m["1"])

	// Check if a key exists in a map
	_, exists := m["3"]
	fmt.Printf("m[3] exists: %v\n", exists)

	// Map with initializer
	n := map[string]int{"a": 1, "b": 2}
	fmt.Println(n)
}
