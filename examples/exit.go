package examples

import (
	"fmt"
	"os"
)

func Exit() {
	defer fmt.Println("!")
	os.Exit(3)
}
