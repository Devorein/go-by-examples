package examples

import (
	"bufio"
	"fmt"
	"net/http"
)

func HttpClients() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	sc := bufio.NewScanner(resp.Body)

	for i := 0; sc.Scan() && i < 5; i++ {
		fmt.Println(sc.Text())
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}
}
