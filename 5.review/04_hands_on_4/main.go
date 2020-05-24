//# Serve the files in the "starting-files" folder

//## To get your images to serve, use:
//``` Go
//	func StripPrefix(prefix string, h Handler) Handler
//	func FileServer(root FileSystem) Handler
//```
//Constraint: you are not allowed to change the route being used for images in the template file

package main

import (
	"log"
	"net/http"
	"text/template"
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
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
