package examples

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumbers() {
	fmt.Println(rand.Intn(100), rand.Intn(100))

	fmt.Println(rand.Float64())
	fmt.Println((rand.Float64()*5)+5, (rand.Float64()*5)+5)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Println(r1.Intn(100), r1.Intn(100))
	r2 := rand.New(rand.NewSource(42))
	fmt.Println(r2.Intn(100), r2.Intn(100), r2.Intn(100), r2.Intn(100))

	r3 := rand.New(rand.NewSource(42))
	fmt.Println(r3.Intn(100), r3.Intn(100))
}
