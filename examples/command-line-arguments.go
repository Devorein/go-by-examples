package examples

import (
	"fmt"
	"os"
)

func CommandLineArguments() {
	args := os.Args
	argsWithoutPath := os.Args[1:]

	fmt.Println(args)
	fmt.Println(argsWithoutPath)
}
