package examples

import (
	"fmt"
	"sort"
)

func Sorting() {
	s := []string{"c", "a", "b"}
	i := []int{3, 1, 2}

	fmt.Println(sort.IntsAreSorted(i))
	fmt.Println(sort.StringsAreSorted(s))

	sort.Strings(s)
	fmt.Println("Sorted strings", s)

	sort.Ints(i)
	fmt.Println("Sorted ints", i)

	fmt.Println(sort.IntsAreSorted(i))
	fmt.Println(sort.StringsAreSorted(s))
}
