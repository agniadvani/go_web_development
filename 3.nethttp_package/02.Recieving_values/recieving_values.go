package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

var t hotdog
var tpl *template.Template

func (t hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		Method        string
		URL           *url.URL
		Submission    map[string][]string
		Header        http.Header
		Host          string
		Contentlength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	err = tpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	err := http.ListenAndServe(":8080", t)
	if err != nil {
		log.Fatal(err)
	}
}
