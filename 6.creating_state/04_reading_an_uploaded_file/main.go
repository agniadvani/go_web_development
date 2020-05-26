package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/form.html"))
}
func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\nfile", f, "\nheader", h)
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
		s = string(bs)
	}
	tpl.ExecuteTemplate(w, "form.html", nil)
	io.WriteString(w, s)
}
