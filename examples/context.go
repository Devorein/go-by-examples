package examples

import (
	"fmt"
	"net/http"
	"time"
)

func hello(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("Hello handler started")
	defer fmt.Println("Hello handler ended")
	fmt.Println(ctx)
	select {
	case <-time.After(time.Second):
		fmt.Fprintf(res, "Hello World")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(res, err.Error(), internalError)
	}
}

func Context() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8089", nil)
}
