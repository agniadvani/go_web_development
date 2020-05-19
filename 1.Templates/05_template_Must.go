package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))

	tpl.ExecuteTemplate(os.Stdout, "boilet.gohtml", nil)

	tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)

	tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)

}
