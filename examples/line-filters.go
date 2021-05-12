package examples

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LineFilters() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		fmt.Println(strings.ToUpper(sc.Text()))
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error", err)
		os.Exit(1)
	}
}
