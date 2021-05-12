package examples

import (
	"fmt"
	"time"
)

func Epoch() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	millis := nanos / 100000
	fmt.Println(now)
	fmt.Println(secs)
	fmt.Println(nanos)
	fmt.Println(millis)
	fmt.Println(time.Unix(secs+25000, 0))
}
