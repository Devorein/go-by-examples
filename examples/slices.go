package examples

import "fmt"

func Slices() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s[0]", s[0])
	fmt.Println("Length of slice", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println(s)

	// Slices are initialized by reference, both c and s points to the same slice
	c1 := s
	c1[2] = "2"
	fmt.Println("s[2]", s[2])

	// Making a copy of slice rather than referencing it
	c2 := make([]string, 5)
	copy(c2, c1)
	fmt.Println(c2)
	c2[2] = "b"
	fmt.Println("s[2]", s[2])

	// Slicing in slices, still references the underlying slice
	l := s[2:]
	l[2] = "5"
	fmt.Println("s[4]", s[4])

	// Initializing slices
	i := [5]int{1, 2, 3, 4, 5}
	fmt.Println(i)

	// dynamic sized 2d slice
	slice := make([][]int, 5)

	for i := range slice {
		// length of inner slices are different every time
		slice[i] = make([]int, i+1)
		for j := range slice[i] {
			slice[i][j] = j
		}
	}

	fmt.Println(slice)
}
