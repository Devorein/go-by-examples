package examples

import (
	b64 "encoding/base64"
	"fmt"
)

func Base64() {
	s := "abc123!@# s"

	sEnc := b64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(sEnc)
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	uEnc := b64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
