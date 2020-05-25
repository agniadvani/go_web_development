package main

import (
	"io"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	tpl := template.Must(template.ParseFiles("templates/form.html"))
	tpl.ExecuteTemplate(w, "form.html", nil)
	io.WriteString(w, v)

}
