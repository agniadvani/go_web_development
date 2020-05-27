package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", counter)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func counter(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("counter")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "counter",
			Value: "0",
		}
	}
	counter, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatal(err)
	}
	counter++
	c.Value = strconv.Itoa(counter)
	http.SetCookie(w, c)
	io.WriteString(w, `<h1>Number of times visited:</h1>`)
	io.WriteString(w, c.Value)
}
