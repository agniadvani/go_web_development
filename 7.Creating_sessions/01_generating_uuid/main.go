package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err == http.ErrNoCookie {
		id, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
		}

		http.SetCookie(w, c)
	}
	fmt.Println(c.Value)
}
