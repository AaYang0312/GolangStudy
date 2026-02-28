package main

import (
	"fmt"
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.String())
	if r.Method != "GET" {
		byteData, _ := io.ReadAll(r.Body)
		fmt.Println(string(byteData))
	}
	fmt.Println(r.Header)

	_, err := w.Write([]byte("hello world"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/index", Index)
	fmt.Println("server start on 127.0.0.1:8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
