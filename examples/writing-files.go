package examples

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WritingFiles() {
	path, err := os.Getwd()
	check(err)
	fp := filepath.Join(filepath.Dir(path), "examples", "test.txt")

	d1 := []byte("\nNew Line")
	err = ioutil.WriteFile(fp, d1, 0644)
	check(err)

	f, err := os.Create(fp)
	check(err)
	defer f.Close()

	n2, err := f.Write(d1)
	check(err)
	fmt.Printf("Wrote %d bytes\n", n2)

	n3, err := f.WriteString("\nAnother line")
	check(err)
	fmt.Printf("Wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.Write(d1)
	check(err)
	fmt.Printf("Wrote %d bytes\n", n4)
	w.Flush()
}
