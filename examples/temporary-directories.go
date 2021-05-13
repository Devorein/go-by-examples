package examples

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func TemporaryDirectories() {
	f, err := ioutil.TempFile("", "sample.txt")
	check(err)

	fmt.Println("Temp file name", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte("Hello World"))
	check(err)

	dname, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("Temp dir name", dname)

	// defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "sample")
	err = ioutil.WriteFile(fname, []byte("Hello World"), 0666)
	check(err)
}
