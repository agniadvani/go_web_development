package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/form.html"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.HandleFunc("/mp", mp)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `<h1>Redirected to foo</h1>`)
	fmt.Println("Request method at foo is:", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request method at bar is:", req.Method)
	//303 see other
	http.Redirect(w, req, "/", 303)
}
func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request method at barred is:", req.Method)
	err := tpl.ExecuteTemplate(w, "form.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func mp(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request method at mp is:", req.Method)
	//moved permanently
	http.Redirect(w, req, "/", 301)
}
