package examples

import (
	"flag"
	"fmt"
)

func CommandLineFlags() {
	strPtr := flag.String("word", "foo", "A string flag")
	intPtr := flag.Int("len", 2, "A int flag")
	boolPtr := flag.Bool("active", false, "A bool flag")

	var strVar string
	flag.StringVar(&strVar, "strVar", "var", "A string variable")

	flag.Parse()

	fmt.Println("word", *strPtr)
	fmt.Println("length", *intPtr)
	fmt.Println("active", *boolPtr)
	fmt.Println("strVar", strVar)
	fmt.Println("Tail", flag.Args())
}
