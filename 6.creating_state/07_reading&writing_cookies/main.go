package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "cookie",
		Value: "42",
	})
	fmt.Fprintln(w, "Cookie written!")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("cookie")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "Value of cookie:")
	fmt.Fprintln(w, c)
}
