package examples

import (
	"fmt"
	"os"
)

func Defer() {
	f := createFile("/temp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(fn string) *os.File {
	fmt.Println("Creating")
	f, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("Closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func writeFile(f *os.File) {
	fmt.Println("Writing")
	fmt.Fprintln(f, "data")
}
