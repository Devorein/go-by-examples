package examples

import (
	"bytes"
	"fmt"
	"regexp"
)

func RegularExpressions() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("p1ach punch"))
	fmt.Println(r.FindStringIndex("p1ach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch", -1))
	fmt.Println(r.FindAllStringIndex("peach punch", -1))
	fmt.Println(r.FindAllStringSubmatch("peach punch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 1))
	fmt.Println(r.ReplaceAllString("a peach with some salt", "fruit"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
