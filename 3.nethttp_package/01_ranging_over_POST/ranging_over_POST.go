package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

var t hotdog

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func (t hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.html", r.Form)
}

func main() {

	err := http.ListenAndServe(":8080", t)
	if err != nil {
		log.Fatalln(err)
	}
}
