package examples

import "os"

func Panic() {
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
