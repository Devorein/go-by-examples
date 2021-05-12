package examples

import (
	"fmt"
	"strconv"
)

func NumberParsing() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(i)

	h, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(h)

	_, err := strconv.ParseUint("123.22", 0, 64)
	fmt.Println(err)

	k, _ := strconv.Atoi("123")
	fmt.Println(k)
}
