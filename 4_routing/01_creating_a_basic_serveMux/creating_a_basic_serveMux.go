package main

import (
	"io"
	"log"
	"net/http"
)

type rockstar int
type game int
type index int

func (i index) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>This is index page</h1>")
}
func (d rockstar) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>ROCKSTARS:</h1><ul><li>Jimmy Page</li><li>David Gilmour</li><li>Keith Richards</li></ul>")
}
func (c game) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>GAME:</h1><ul><li>Super Mario</li><li>Pacman</li><li>Call of Duty</li></ul>")
}

func main() {
	var i index
	var c game
	var d rockstar
	mux := http.NewServeMux()
	mux.Handle("/", i)
	mux.Handle("/game", c)
	mux.Handle("/rockstar", d)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
