package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello \nplease Enter (http://localhost:8090/goByExamples)!"))
}

func headers(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadFile("index/index.html")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(bytes))
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/goByExamples", headers)
	http.ListenAndServe(":8090", nil)
}
