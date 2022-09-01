package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello \nplease Enter (https://localhost:8090/goByExamples)!"))
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/goByExamples", headers)
	http.ListenAndServe(":8090", nil)
}
