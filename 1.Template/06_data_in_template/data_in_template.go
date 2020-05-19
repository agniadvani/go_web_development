package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))

	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", `When the going gets tough,the tough gets going.`)
	if err != nil {
		log.Fatal(err)
	}
}
