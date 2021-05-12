package examples

import (
	"fmt"
	"sort"
)

type byLength []string

func (b byLength) Len() int {
	return len(b)
}

func (b byLength) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byLength) Less(i, j int) bool {
	return len(b[i]) < len(b[j])
}

func SortingByFunctions() {
	s := []string{"apple", "orange", "kiwi"}
	sort.Sort(byLength(s))
	fmt.Println(s)
}
