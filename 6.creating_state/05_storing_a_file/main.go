package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

		//open the file
		f, h, err := req.FormFile("q")
		if err != nil {
			log.Fatal(err)
		}
		//read the file
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
		s = string(bs)

		//store the file
		fmt.Println(h.Filename)
		dst, err := os.Create(filepath.Join("../../", h.Filename))
		if err != nil {
			log.Fatal(err)
		}
		defer dst.Close()
		io.WriteString(dst, s)
	}
	err := tpl.ExecuteTemplate(w, "form.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
