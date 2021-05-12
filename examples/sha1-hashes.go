package examples

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func Sha1Hashes() {
	s := "sha1 this string"

	md5h := md5.New()

	md5h.Write([]byte(s))
	md5bs := md5h.Sum(nil)
	fmt.Printf("%x\n", md5bs)

	sha1h := sha1.New()
	sha1h.Write([]byte(s))
	fmt.Printf("%x\n", sha1h.Sum(nil))
}
