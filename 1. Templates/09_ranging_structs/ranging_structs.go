package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type rockstar struct {
	Name string
	Band string
}

func main() {
	tpl = template.Must(template.ParseGlob("template/template.gohtml"))

	r1 := rockstar{
		`David Gilmour`,
		`Pink Floyd`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", r1)
	if err != nil {
		log.Fatal(err)
	}
}
