package examples

import (
	"fmt"
	"net/http"
)

func Hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello\n")
}

func Header(res http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(res, "%v,%v\n", name, h)
		}
	}
}

func HttpServers() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/header", Header)

	http.ListenAndServe(":8090", nil)
}
