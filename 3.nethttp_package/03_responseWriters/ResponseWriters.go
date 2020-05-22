package main

import (
	"io"
	"log"
	"net/http"
)

type hotdog int

func (t hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Rock-Type", "Roger Waters")
	w.Header().Set("Content-Type", "text/html, charset = utf-8")
	io.WriteString(w, `<h1>HTML</h1>`)
}

func main() {
	var d hotdog
	err := http.ListenAndServe(":8080", d)
	if err != nil {
		log.Fatal(err)
	}
}
