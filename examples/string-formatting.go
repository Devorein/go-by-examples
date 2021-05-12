package examples

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func StringFormatting() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", false)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 3)
	fmt.Printf("%c\n", 68)
	fmt.Printf("%x\n", 554)
	fmt.Printf("%f\n", 25.9)
	fmt.Printf("%e\n", 25123.9123213)
	fmt.Printf("%E\n", 25123.9123213)
	fmt.Printf("%s\n", "\"String\"")
	fmt.Printf("%q\n", "\"String\"")
	fmt.Printf("%x\n", "hex this")
	fmt.Printf("%p\n", &p)
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 12.231, 345.1)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 12.231, 345.1)
	fmt.Printf("|%6s|%6s|\n", "av", "foobar")
	fmt.Printf("|%-6s|%-6s|\n", "av", "foobar")
	s := fmt.Sprintf("%d\n", 123)
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "%s", "error")
}
