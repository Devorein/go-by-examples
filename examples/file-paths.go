package examples

import (
	"fmt"
	"path/filepath"
	"strings"
)

func FilePaths() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println(p)

	fmt.Println(filepath.Join("dir1//", "dir2"))
	fmt.Println(filepath.Join("dir1/../dir2", "filename"))

	fmt.Println(filepath.Dir(p))
	fmt.Println(filepath.Base(p))

	fmt.Println(filepath.IsAbs("\\dir1\\dir2"))
	fmt.Println(filepath.IsAbs("dir1\\dir2"))

	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/c/d")
	check(err)
	fmt.Println(rel)

	rel, _ = filepath.Rel("a/b", "a/c/d")
	fmt.Println(rel)
}
