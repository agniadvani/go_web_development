package main

import (
	"io"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>This is index page</h1>")
}
func rockstar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>ROCKSTARS:</h1><ul><li>Jimmy Page</li><li>David Gilmour</li><li>Keith Richards</li></ul>")
}
func game(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>GAME:</h1><ul><li>Super Mario</li><li>Pacman</li><li>Call of Duty</li></ul>")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/game", game)
	http.HandleFunc("/rockstar", rockstar)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
