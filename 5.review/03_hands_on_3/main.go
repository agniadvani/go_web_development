//Serve the files in the "starting-files" folder
//To get your images to serve, use only this:
//	fs := http.FileServer(http.Dir("public"))
//Hint: look to see what type FileServer returns, then think it through.

package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}
func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
