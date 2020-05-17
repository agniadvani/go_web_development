package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("template/template.gohtml"))
	rockstars := []string{`David Gilmour`, `Jimmy Page`, `Roger Waters`, `Mick Jagger`}
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", rockstars)
	if err != nil {
		log.Fatal(err)
	}
}
