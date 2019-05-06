package main

import (
	"net/http"
	"log"
	"fmt"
	"gopl.io/ch1/lissajous"
)

func main() {
	//http.HandleFunc("/", handler) // each request calls handler

	//打印动画
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous(w, r)
	}

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
