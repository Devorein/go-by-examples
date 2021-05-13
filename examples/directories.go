package examples

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Directories() {
	j := filepath.Join
	path, err := os.Getwd()
	check(err)
	dp := j(filepath.Dir(path), "examples", "nested")
	check(os.Mkdir(dp, 0755))
	defer os.RemoveAll(dp)

	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile(j(dp, "test.txt"))
	check(os.MkdirAll(j(dp, "nested2", "nested3"), 0755))
	createEmptyFile(j(dp, "nested2", "test.txt"))
	createEmptyFile(j(dp, "nested2", "nested3", "test.txt"))

	c, err := ioutil.ReadDir(j(dp))
	check(err)
	for _, entry := range c {
		fmt.Println(entry.Name(), entry.IsDir())
	}

	check(os.Chdir(j(dp, "nested2")))

	c, err = ioutil.ReadDir(".")
	check(err)
	for _, entry := range c {
		fmt.Println(entry.Name(), entry.IsDir())
	}

	check(filepath.Walk(dp, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path, info.IsDir())
		return nil
	}))
}
