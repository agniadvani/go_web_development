package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "cookie",
		Value: "42",
	})
	io.WriteString(w, `<h1>Cookie has been added</h1><br><br>`)
	io.WriteString(w, `<a href="/read"><em>Read Cookie</em></a>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("cookie")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	fmt.Fprintln(w, `<h1>Cookie value:</h1><br>`, c)
	fmt.Fprintln(w, `<br> <a href="/expire">Delete cookie</a>`)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("cookie")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1 //deleting a cookie
	http.SetCookie(w, c)
	fmt.Fprintln(w, "<h1>Cookie Deleted</h1><br><br>")
	fmt.Fprintln(w, ` <a href="/">Set a Cookie</a>`)
}
