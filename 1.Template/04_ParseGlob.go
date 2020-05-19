package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("template/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(os.Stdout, "boilet.gohtml", nil)

	tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)

	tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)

}
